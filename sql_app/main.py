from fastapi import Depends, FastAPI
import uvloop
from sqlalchemy.orm import Session

from . import models, schemas
from sql_app import crud
from sql_app.database import engine, Base, AsyncSession, async_session

uvloop.install()

async def init_models():
    async with engine.begin() as conn:
        await conn.run_sync(Base.metadata.drop_all)
        await conn.run_sync(Base.metadata.create_all)

app = FastAPI()


# Dependency
async def get_session() -> AsyncSession:
    async with async_session() as session:
        yield session


@app.post("/users/", response_model=schemas.UserOut)
async def create_user(user: schemas.UserCreate, db: AsyncSession = Depends(get_session)):
    return await crud.create_user(db=db, user=user)


@app.get("/users/", response_model=list[schemas.UserOut])
async def read_users(skip: int = 0, limit: int = 100, db: AsyncSession = Depends(get_session)):
    users = await crud.get_users(db, skip=skip, limit=limit)
    return users
