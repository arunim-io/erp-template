import strawberry
from accounts.mutations import AuthMutations
from accounts.queries import AuthQueries
from strawberry_django.optimizer import DjangoOptimizerExtension


@strawberry.type
class Query:
    @strawberry.field
    def auth(self) -> AuthQueries:
        return AuthQueries()


@strawberry.type
class Mutation:
    @strawberry.field
    def auth(self) -> AuthMutations:
        return AuthMutations()


schema = strawberry.Schema(
    query=Query,
    mutation=Mutation,
    extensions=[DjangoOptimizerExtension],
)
