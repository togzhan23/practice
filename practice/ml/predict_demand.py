import pandas as pd
import numpy as np
from sklearn.linear_model import LinearRegression
from sklearn.preprocessing import StandardScaler
import joblib
import os

def train_model():
    data = pd.DataFrame({
        "month": np.random.randint(1, 13, 100),
        "price_per_ton": np.random.uniform(20000, 50000, 100),
        "stock": np.random.uniform(100, 1000, 100),
        "demand": np.random.uniform(50, 800, 100)
    })

    X = data[["month", "price_per_ton", "stock"]]
    y = data["demand"]

    scaler = StandardScaler()
    X_scaled = scaler.fit_transform(X)

    model = LinearRegression()
    model.fit(X_scaled, y)

    # Сохраняем модель и скейлер
    os.makedirs("ml/models", exist_ok=True)
    joblib.dump(model, "ml/models/demand_model.pkl")
    joblib.dump(scaler, "ml/models/scaler.pkl")


def predict_demand(month: int, price_per_ton: float, stock: float) -> float:
    model = joblib.load("ml/models/demand_model.pkl")
    scaler = joblib.load("ml/models/scaler.pkl")

    X = pd.DataFrame([[month, price_per_ton, stock]], columns=["month", "price_per_ton", "stock"])
    X_scaled = scaler.transform(X)

    prediction = model.predict(X_scaled)
    return round(prediction[0], 2)
