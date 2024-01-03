import tensorflow as tf
import numpy as np
from tensorflow.keras.layers import *
from keras_cv_attention_models import edgenext
from tensorflow.keras import Model 

def get_base_model(name, input_shape):
    if name == 'EdgeNeXt_XX_Small':
        return edgenext.EdgeNeXt_XX_Small(num_classes=0, input_shape=input_shape, pretrained="imagenet")

    raise Exception("Cannot find this base model:", name)

def get_out_layers(name):
    if name == 'EdgeNeXt_XX_Small':
        return [
            'stack2_downsample_ln',
            'stack3_downsample_ln',
            'stack4_downsample_ln',
            'stack4_block2_stda_output'
        ]
    
    return None

def up_sample(inputs, steps=2):
    '''
    Upsample the input by a factor of 2
    '''
    x = UpSampling2D((steps, steps))(inputs)
    return x

def conv(inputs, filters, kernel_size=3, strides=(1, 1), padding='same', activation=None):
    x = Conv2D(filters, kernel_size=kernel_size, strides=strides, padding=padding, activation=activation)(inputs)
    return x

def bn_act(inputs, activation='swish'):
    x = BatchNormalization()(inputs)
    if activation is not None:
        x = Activation(activation)(x)
    return x

def conv_bn_act(inputs, filters, kernel_size, strides=(1, 1), padding='same', activation='swish'):
    x = Conv2D(filters, kernel_size=kernel_size, strides=strides, padding=padding)(inputs)
    x = bn_act(x, activation=activation)
    return x

def residual_block(inputs, filters):
    x = conv_bn_act(inputs, filters, 3)
    x = conv(x, filters, 3)
    x = BatchNormalization()(x)
    skips = conv(inputs, filters, 1) if inputs.shape[-1] != filters else inputs
    x = x + skips
    x = Activation('swish')(x)
    return x

def create_model(input_shape=(352,352,3), num_classes=1, base_name='EdgeNeXt_XX_Small'):
    base = get_base_model(base_name, input_shape)

    backbone_layer_names = get_out_layers(base_name)
    backbone_layers = [base.get_layer(layer_name).output for layer_name in backbone_layer_names] # big to small
    backbone_layers.reverse() # small to big

    filters = 128

    x = backbone_layers[0]

    for layer in backbone_layers[1:]:
        x = up_sample(x, 2)
        x = Concatenate()([x, layer])
        x = Dropout(0.01)(x)
        x = residual_block(x, filters)

    x = up_sample(x, 4)
    x = Dropout(0.01)(x)
    x = residual_block(x, filters)

    x = Conv2D(num_classes, 1)(x)
    outputs = Activation('sigmoid')(x)

    model = Model(base.input, outputs)
    
    return model