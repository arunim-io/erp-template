import strawberry
from accounts.mutations import AuthMutation
from accounts.queries import AuthQueries
from strawberry.tools import merge_types
from strawberry_django.optimizer import DjangoOptimizerExtension

schema = strawberry.Schema(
    query=merge_types("Query", (AuthQueries,)),
    mutation=merge_types("Mutation", (AuthMutation,)),
    extensions=[DjangoOptimizerExtension],
)
