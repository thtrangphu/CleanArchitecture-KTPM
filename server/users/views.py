from django.contrib.auth.hashers import make_password
from django.contrib.auth.models import User
from django.http import JsonResponse
from rest_framework import status
from rest_framework.permissions import AllowAny, IsAuthenticated
from rest_framework.response import Response
from rest_framework.utils import json
from rest_framework.views import APIView

from .serializers import UserSerializer, ChangePasswordSerializer, UpdateProfileSerializer


class UserView(APIView):
    queryset = User.objects.all()
    serializer_class = UserSerializer
    permission_classes = (AllowAny,)

    def post(self, request, *args, **kwargs):
        serializer = UserSerializer(data=request.data)
        if serializer.is_valid():
            user = serializer.save()

            return JsonResponse({
                'message': 'Register successful! Please login again.'
            }, status=status.HTTP_201_CREATED)

        else:
            return JsonResponse({
                'error_message': serializer.errors,
                'errors_code': 400,
            }, status=status.HTTP_400_BAD_REQUEST)


class UserProfileView(APIView):
    serializer_class = UserSerializer
    permission_classes = (IsAuthenticated,)

    def get(self, request, *args, **kwargs):
        return Response({
            'id': request.user.id,
            'username': request.user.username,
            'email': request.user.email,
            'firstName': request.user.first_name,
            'lastName': request.user.last_name
        })

    def put(self, request, *args, **kwargs):
        serializer = UpdateProfileSerializer(request.user, data=request.data, partial=True)
        serializer.is_valid(raise_exception=True)
        return Response(status=status.HTTP_200_OK)


    def patch(self, request, *args, **kwargs):
        serializer = ChangePasswordSerializer(data=request.data)
        if serializer.is_valid():
            # Check old password
            if not request.user.check_password(serializer.data.get("old_password")):
                return Response({"old_password": ["Wrong password."]}, status=status.HTTP_400_BAD_REQUEST)
            # set_password also hashes the password that the user will get
            request.user.set_password(serializer.data.get("new_password"))
            request.user.save()
            response = {
                'status': 'success',
                'code': status.HTTP_200_OK,
                'message': 'Password updated successfully',
                'data': []
            }

            return Response(response)

        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)
