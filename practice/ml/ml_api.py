from fastapi import FastAPI
from pydantic import BaseModel
from predict_demand import predict_demand

app = FastAPI(title="ML API for Demand Forecasting")

class DemandRequest(BaseModel):
    month: int
    price_per_ton: float
    stock: float

@app.post("/predict")
def get_demand_forecast(data: DemandRequest):
    prediction = predict_demand(data.month, data.price_per_ton, data.stock)
    return {
        "predicted_demand": prediction,
        "unit": "тонн"
    }
