import cv2
import numpy as np

from model import create_model

# load model

model_path = "polyp_model_256_1_class.h5"

im_size = 256
num_classes = 1

model = create_model(input_shape=(im_size, im_size, 3), num_classes=num_classes, base_name='EdgeNeXt_XX_Small')

model.load_weights(model_path)

# how to use

img_path = "./images/109.png"

img = cv2.imread(img_path)
img = cv2.cvtColor(img, cv2.COLOR_BGR2RGB)
img = cv2.resize(img, (im_size, im_size))
img = img / 255.0

preds = model.predict(np.expand_dims(img, 0), verbose=0)
pred = np.squeeze(preds)
pred = np.round(pred)
pred = pred * 255

cv2.imwrite("pred2.png", pred)