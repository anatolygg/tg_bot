from fastapi import FastAPI
from pydantic import BaseModel
from ml_service import find_answer

app = FastAPI()

class QuestionRequest(BaseModel):
    question: str

@app.post("/predict")
def predict(request: QuestionRequest):
    question = request.question
    answer = find_answer(question)
    
    response = answer[0]
    return {"answer": response}