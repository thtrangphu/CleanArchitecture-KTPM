# Generated by Django 5.0 on 2023-12-17 16:51

import django.db.models.deletion
from django.conf import settings
from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        migrations.swappable_dependency(settings.AUTH_USER_MODEL),
    ]

    operations = [
        migrations.CreateModel(
            name='Recognition',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('raw', models.FileField(upload_to='')),
                ('result', models.FileField(upload_to='')),
                ('uploaded_on', models.DateTimeField(auto_now_add=True)),
                ('owner', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='recognitions', to=settings.AUTH_USER_MODEL)),
            ],
        ),
    ]
