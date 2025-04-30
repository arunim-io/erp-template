import strawberry_django

from . import models


@strawberry_django.type(
    models.User,
    fields=["id", "username", "email", "is_active", "first_name", "last_name"],
)
class User: ...
