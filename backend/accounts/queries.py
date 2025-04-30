import strawberry
import strawberry_django

from . import types


@strawberry.type
class AuthQueries:
    status: types.User = strawberry_django.auth.current_user()
