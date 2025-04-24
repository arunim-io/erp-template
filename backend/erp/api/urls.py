from django.urls import include, path
from drf_spectacular.views import SpectacularAPIView, SpectacularSwaggerView

urlpatterns = [
    path("schema", SpectacularAPIView.as_view(), name="schema"),
    path("", SpectacularSwaggerView.as_view(), name="schema-docs"),
    path("accounts/", include("accounts.api.urls")),
]
