"""
URL configuration for djangoProject project.

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/5.0/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls '))
"""
from django.conf.urls.static import static
from django.contrib import admin
from django.urls import path
from rest_framework_simplejwt.views import (
    TokenObtainPairView,
    TokenRefreshView,
)

from djangoProject import settings
from recognitions.views import RecognitionView, RecognitionDetailView
from users.views import UserView, UserProfileView

urlpatterns = [
    path('admin/', admin.site.urls),
    path('api/auth/', TokenObtainPairView.as_view(), name='token_obtain_pair'),
    path('api/auth/refresh/', TokenRefreshView.as_view(), name='token_refresh'),
    path('api/users/', UserView.as_view(), name='api_user'),
    path('api/users/profile/', UserProfileView.as_view(), name='api_user_profile'),
    path('api/recognitions/', RecognitionView.as_view(), name='api_recognition'),
    path('api/recognitions/<int:pk>', RecognitionDetailView.as_view(), name='api_recognition_detail')
] + static(settings.MEDIA_URL, document_root=settings.MEDIA_ROOT)
