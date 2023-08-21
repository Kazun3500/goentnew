from sqlalchemy.orm import Session
from sqlalchemy import select

from . import models, schemas
from .database import AsyncSession


async def get_users(db: AsyncSession, skip: int = 0, limit: int = 100):
    result =  await db.execute(select(models.User).offset(skip).limit(limit).all())
    return result.scalars().all()


async def create_user(db: AsyncSession, user: schemas.UserCreate):
    db_user = models.User(name=user.name, age=user.age)
    db.add(db_user)
    await db.commit()
    return db_user
