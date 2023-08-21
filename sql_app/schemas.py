from pydantic import BaseModel, PositiveInt


class UserCreate(BaseModel):
    name: str
    age: PositiveInt


class UserOut(BaseModel):
    id: int
    name: str
    age: PositiveInt

    class Config:
        orm_mode = True