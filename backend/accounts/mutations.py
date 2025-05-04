import strawberry
import strawberry_django

from . import types


@strawberry.type
class AuthMutations:
    login: types.User = strawberry_django.auth.login()
    logout = strawberry_django.auth.logout()
