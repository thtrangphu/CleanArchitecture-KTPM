from django.shortcuts import render
from django.views import View
from rest_framework import generics, status
from rest_framework.parsers import FormParser, MultiPartParser
from rest_framework.permissions import AllowAny, IsAuthenticated
from rest_framework.response import Response
from rest_framework.views import APIView

from recognitions.models import Recognition
from recognitions.serializer import RecognitionSerializer


class RecognitionView(APIView):
    parser_classes = (MultiPartParser, FormParser)
    permission_classes = (IsAuthenticated,)
    serializer_class = RecognitionSerializer

    def get(self, request, *args, **kwargs):
        history = Recognition.objects.filter(owner_id=request.user.id)
        serializer = self.serializer_class(history, many=True)
        print(history)
        return Response({"message": "Recognition history", "history": serializer.data},
                        status=status.HTTP_200_OK)

    def post(self, request, *args, **kwargs):
        serializer = self.serializer_class(data=request.data)
        if serializer.is_valid():
            serializer.save(owner=self.request.user)
            return Response(status=status.HTTP_201_CREATED)

        return Response(
            serializer.errors,
            status=status.HTTP_400_BAD_REQUEST
        )


class RecognitionDetailView(APIView):
    permission_classes = (IsAuthenticated,)
    serializer_class = RecognitionSerializer

    def get(self, request, *args, **kwargs):
        record = Recognition.objects.get(id=self.kwargs['pk'])
        serializer = self.serializer_class(record)
        return Response({'message': 'Recognition detail', 'recognition': serializer.data},
                        status=status.HTTP_200_OK)
