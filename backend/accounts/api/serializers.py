from django.contrib.auth import authenticate
from django.utils.translation import gettext_lazy as _
from knox.models import AuthToken
from rest_framework_json_api import serializers

from accounts.models import User


class UserSerializer(serializers.ModelSerializer):
    class Meta:
        resource_name = "auth"
        model = User
        fields = [User.USERNAME_FIELD, User.EMAIL_FIELD, "first_name", "last_name"]


class LoginSerializer(serializers.Serializer):
    username = serializers.CharField(label=_("Username"), write_only=True)
    password = serializers.CharField(
        label=_("Password"),
        style={"input_type": "password"},
        trim_whitespace=False,
        write_only=True,
    )

    class Meta:
        resource_name = "auth"

    def validate(self, attrs):
        username, password = attrs.get("username"), attrs.get("password")

        if username and password:
            user = authenticate(
                request=self.context.get("request"),
                username=username,
                password=password,
            )

            if not user:
                self.raise_validation_error(
                    "Unable to log in with provided credentials."
                )
        else:
            self.raise_validation_error('Must include "username" and "password".')

        attrs["user"] = user
        return attrs

    def raise_validation_error(self, msg: str):
        raise serializers.ValidationError(_(msg), "authorization")


class LoginResponseSerializer(serializers.ModelSerializer):
    id = serializers.CharField(source="token_key")
    token_expiry_date = serializers.DateTimeField(source="expired_at", read_only=True)
    user = UserSerializer(read_only=True)
    token = serializers.SerializerMethodField()

    class Meta:
        resource_name = "auth"
        model = AuthToken
        fields = ["id", "token_expiry_date", "user", "token"]

    def get_token(self, obj: AuthToken) -> str:
        return str(self.context["token"])
