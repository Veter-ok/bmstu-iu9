import joblib
import os

# Пути к файлам (предполагается, что они лежат в корне проекта)
BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
VECTORIZER_PATH = os.path.join(BASE_DIR, "ml/vectorizer.pkl")
MODEL_PATH = os.path.join(BASE_DIR, "ml/sentiment_model.pkl")

# Загружаем при старте (один раз)
vectorizer = joblib.load(VECTORIZER_PATH)
sentiment_model = joblib.load(MODEL_PATH)

import json
from django.contrib.auth.mixins import LoginRequiredMixin
from django.core.exceptions import ValidationError
from django.http import JsonResponse
from django.shortcuts import render, redirect
from django.views import View
from .models import UserToken, Review

def about(request):
    return render(request, 'myapp/about.html')

def contacts(request):
    return render(request, 'myapp/contacts.html')

def images(request):
    return render(request, 'myapp/images.html')

class DeleteTokenView(LoginRequiredMixin, View):
    def post(self, request, token_id):
        try:
            token = UserToken.objects.get(id=token_id, user=request.user)
            token.delete()
            return JsonResponse({"status": "ok"})
        except UserToken.DoesNotExist:
            return JsonResponse({"status": "error", "message": "Token not found"}, status=404)


class DeleteReviewView(LoginRequiredMixin, View):
    def post(self, request, review_id):
        try:
            review = Review.objects.get(id=review_id, user=request.user)
            review.delete()
            return JsonResponse({"status": "ok"})
        except Review.DoesNotExist:
            return JsonResponse({"status": "error", "message": "Review not found"}, status=404)

from django.utils.functional import cached_property

class AddReviewView(View):
    def post(self, request):
        text = request.POST.get("text", "")
        sentiment = self.classify_text(text)

        Review.objects.create(
            text=text,
            sentiment=sentiment,
            user=request.user
        )

        return redirect("reviews")

    def classify_text(self, text):
        X = vectorizer.transform([text])
        prediction = sentiment_model.predict(X)[0]
        return bool(prediction) 

class ReviewPage(LoginRequiredMixin, View):
    def get(self, request):
        reviews = Review.objects.all().order_by("-id")
        return render(request, "myapp/reviews.html", {"reviews": reviews})

class Index(LoginRequiredMixin, View):
    def get(self, request):
        return redirect("reviews")
        user = request.user
        user_tokens = user.tokens.all()
        return render(request, "myapp/index.html", {"user_tokens": user_tokens})
    def post(self, request):
        data = json.loads(request.body.decode("utf-8"))
        if "token_name" not in data.keys():
            JsonResponse({"errors": "Bad request"}, status=400)
        token_name = data["token_name"]
        user = request.user
        new_token = UserToken(
            user=user,
            name=token_name,
        )
        try:
            new_token.full_clean()
            new_token.save()
            return JsonResponse(
                {
                    "token_name": new_token.name,
                    "token": new_token.token,
                },
                status=200,
            )
        except ValidationError as e:
            return JsonResponse({"detail": e.messages}, status=400)