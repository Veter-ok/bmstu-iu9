from django.urls import path
from django.contrib.auth.views import LoginView, LogoutView
from . import views

urlpatterns = [
    path("login/", LoginView.as_view(), name="login"),
    path("logout/", LogoutView.as_view(), name="logout"),
    path("", views.Index.as_view(), name="index"),  # генерация токенов
    path("reviews/", views.ReviewPage.as_view(), name="reviews"),  # страница отзывов
    path("add-review/", views.AddReviewView.as_view(), name="add-review"),
    path("delete-review/<int:review_id>/", views.DeleteReviewView.as_view(), name="delete-review"),
    path('about/', views.about, name='about'),
    path('images/', views.images, name='images'),
    path('contacts/', views.contacts, name='contacts'),
]