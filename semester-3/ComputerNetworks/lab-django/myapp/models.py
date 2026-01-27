from django.db import models
from django.contrib.auth.models import User
from django.core.validators import MaxLengthValidator
import secrets

class Review(models.Model):
    user = models.ForeignKey(User, on_delete=models.CASCADE, related_name="reviews")
    text = models.TextField()
    sentiment = models.BooleanField()  # True = позитивный, False = негативный
    created_at = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return f"{self.user.username} — {self.text[:20]}"

class UserToken(models.Model):
    user = models.ForeignKey(
        User,
        editable=False,
        on_delete=models.CASCADE,
        related_name="tokens",
        verbose_name="Пользователь",
    )
    name = models.CharField(
        max_length=100,
        verbose_name="Название токена",
        validators=[MaxLengthValidator(100)],
    )
    token = models.CharField(
        max_length=64, unique=True, editable=False, verbose_name="Значение токена"
    )
    created_at = models.DateTimeField(auto_now_add=True, verbose_name="Дата создания")


    class Meta:
        verbose_name = "Токен пользователя"
        verbose_name_plural = "Токены пользователей"
        ordering = ["-created_at"]
        unique_together = [["user", "name"]]

    def __str__(self):
        return f"{self.name} ({self.user.username})"

    def save(self, *args, **kwargs):
        if not self.token:
            self.token = self.generate_token()
        super().save(*args, **kwargs)

    @staticmethod
    def generate_token(length=32):
        return secrets.token_hex(length)