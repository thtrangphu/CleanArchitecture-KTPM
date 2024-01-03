from django.contrib.auth.models import User
from django.db import models


class Recognition(models.Model):
    raw = models.FileField()
    result = models.FileField()
    uploaded_on = models.DateTimeField(auto_now_add=True)
    owner = models.ForeignKey(
        User, on_delete=models.CASCADE, related_name="recognitions")
