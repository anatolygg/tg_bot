import os
import pandas as pd
import re
from sentence_transformers import SentenceTransformer, util


os.environ["WANDB_MODE"] = "disabled"


# Загрузка данных из файла Excel
train_data = pd.read_excel('FAQ.xlsx')

# Предобработка данных
def preprocess_text(text):
    text = text.lower()  # Приведение текста к нижнему регистру
    text = re.sub(r'[^а-яА-Я\s]', '', text)  # Удаление символов, кроме букв и пробелов
    return text.strip()

# Удаление пустых значений
train_data = train_data.dropna(subset=['Вопрос', 'Ответ'])

# Применение предобработки к вопросам и ответам
train_data['Вопрос'] = train_data['Вопрос'].apply(preprocess_text)
train_data['Ответ'] = train_data['Ответ']

questions_train = train_data['Вопрос'].tolist()
answers_train = train_data['Ответ'].tolist()

# Загрузка модели SBERT
model = SentenceTransformer('distiluse-base-multilingual-cased-v2')

# Создание эмбеддингов для всех вопросов в обучающей выборке
question_embeddings = model.encode(questions_train, convert_to_tensor=True)

# Функция для нахождения ответа на вопрос
def find_answer(user_question, threshold=0.5):
    user_question_embedding = model.encode(user_question, convert_to_tensor=True)
    similarities = util.pytorch_cos_sim(user_question_embedding, question_embeddings)[0]
    best_match_idx = similarities.argmax().item()
    best_match_score = similarities[best_match_idx].item()

    if best_match_score < threshold:
        return "Я еще не знаю ответы на все вопросы, но я знаю у кого они есть!\nОбратитесь в справочную службу МИФИ.\nНужный номер вы можете найти на сайте: https://mephi.ru/distance-edu/for-students/hot-line.", best_match_score  # Возвращать сообщение о том, что ответ не найден
    
    print(answers_train[best_match_idx])
    return answers_train[best_match_idx], best_match_score