from typing import override

from django.contrib.auth import login, logout
from drf_spectacular.utils import extend_schema
from knox import views as knox_views
from knox.auth import TokenAuthentication
from knox.models import AuthToken
from rest_framework import permissions, status
from rest_framework.request import Request
from rest_framework.response import Response
from rest_framework.views import APIView

from . import serializers


class LoginView(knox_views.LoginView):
    resource_name = "auth"
    permission_classes = [permissions.AllowAny]

    @extend_schema(
        request=serializers.LoginRequestSerializer,
        responses={
            200: serializers.LoginResponseSerializer,
        },
    )
    @override
    def post(self, request: Request, format=None):
        serializer = serializers.LoginRequestSerializer(data=request.data)
        serializer.is_valid(raise_exception=True)

        user = serializer.validated_data["user"]
        login(request, user)

        return super().post(request, format)

    @override
    def get_post_response_data(self, request: Request, token: str, instance: AuthToken):
        return serializers.LoginResponseSerializer(
            instance,
            context={"token": token},
        ).data


class LogoutView(APIView):
    """Logs out the user & remove their related auth token."""

    resource_name = "auth"
    authentication_classes = [TokenAuthentication]
    permission_classes = [permissions.IsAuthenticated]

    @extend_schema(responses={204: None})
    def delete(self, request: Request, format=None):
        request._auth.delete()  # noqa: SLF001

        logout(request)

        return Response(None, status=status.HTTP_204_NO_CONTENT)
