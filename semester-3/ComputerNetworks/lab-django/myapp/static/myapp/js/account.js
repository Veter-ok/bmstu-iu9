function token_copy() {
  const token = this.parentElement.querySelector(".token-value").textContent;
  navigator.clipboard.writeText(token).then(() => {
    const originalText = this.textContent;
    this.textContent = "Скопировано!";
    setTimeout(() => {
      this.textContent = originalText;
    }, 2000);
  });
}

function deleteItem() {
  const id = this.dataset.id;
  const url = `/delete-token/${id}/`; // или `/delete-review/${id}/` для отзывов

  fetch(url, {
    method: "POST",
    headers: {
      "X-CSRFToken": token
    }
  })
    .then(response => response.json())
    .then(data => {
      if (data.status === "ok") {
        this.closest(".token").remove(); // удаляем элемент со страницы
      } else {
        alert("Ошибка при удалении: " + data.message);
      }
    });
}

// Назначаем обработчик
document.querySelectorAll(".del-btn").forEach(button => {
  button.addEventListener("click", deleteItem);
});

// Функция для копирования токенов
document.querySelectorAll(".copy-btn").forEach((button) => {
  button.addEventListener("click", token_copy);
});
// Валидация форм
document
  .getElementById("apply-btn")
  .addEventListener("click", validateFirstForm);
const firstTokenName = document.getElementById("token-name");
const firstTokenDays = document.getElementById("days");

function activateTokenNameError() {
  document.getElementById("token-select-error").classList.add("visible");
  document.getElementById("token-select-error").innerHTML =
    "Пожалуйста, введите уникальное имя (не более 100 символов)";
}
function deactivateTokenNameError() {
  document.getElementById("token-select-error").classList.remove("visible");
  document.getElementById("token-select-error").innerHTML = "Пусто";
}

function validateFirstForm() {
  let isValid = true;

  if (!firstTokenName.value || firstTokenName.value.length > 100) {
    activateTokenNameError();
    isValid = false;
  } else {
    deactivateTokenNameError();
  }

  if (isValid) {
    sendFirstForm();
  }
}
function response_check(response) {
  if (!response.ok) {
    console.log(response.statusText);
    activateTokenNameError();
    throw new Error(response.statusText);
  }
  console.log(response.json());
}

function sendFirstForm() {
  fetch(index, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "X-CSRFToken": token,
    },
    body: JSON.stringify({
      token_name: firstTokenName.value,
    }),
  })
    .then(response_check)
    .catch((error) => console.error("Ошибка:", error));
}
