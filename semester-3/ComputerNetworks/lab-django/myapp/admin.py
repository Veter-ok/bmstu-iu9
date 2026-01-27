from django.contrib import admin
from .models import UserToken

@admin.register(UserToken)
class TwoGisTokenAdmin(admin.ModelAdmin):
    list_display = ('user', 'name', 'created_at')