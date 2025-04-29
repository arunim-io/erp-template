import strawberry
from strawberry_django.optimizer import DjangoOptimizerExtension


@strawberry.type
class Query:
    @strawberry.field
    def config(self) -> str:
        return "WIP"


schema = strawberry.Schema(
    query=Query,
    extensions=[DjangoOptimizerExtension],
)
