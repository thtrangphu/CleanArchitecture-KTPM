from rest_framework import serializers

from recognitions.models import Recognition


class RecognitionSerializer(serializers.ModelSerializer):
    query_set = Recognition.objects.all()
    result = serializers.FileField(required=False)
    raw = serializers.FileField( required=True)
    owner = serializers.ReadOnlyField(source='owner.username')
    owner_id = serializers.ReadOnlyField(source='owner.id')

    class Meta:
        model = Recognition
        fields = ['raw', 'result', 'uploaded_on', 'owner', 'owner_id']


