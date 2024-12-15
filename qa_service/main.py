from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()

class QuestionRequest(BaseModel):
    question: str

@app.post("/predict")
def predict(request: QuestionRequest):
    request = QuestionRequest("How are u?")
    
    question = request.question
    response = f"Ответ на вопрос '{question}' связан с МИФИ."
    return {"answer": response}