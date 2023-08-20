from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.ext.asyncio import create_async_engine
from sqlalchemy.orm import sessionmaker

# SQLALCHEMY_DATABASE_URL = "sqlite:///./sql_app.db"
# "host=127.0.0.1 port=5434 user=postgres dbname=postgres password=mysecretpassword sslmode=disable"
SQLALCHEMY_DATABASE_URL = "postgresql+asyncpg://postgres:mysecretpassword@127.0.0.1:5434/postgres"


engine = create_async_engine(SQLALCHEMY_DATABASE_URL, echo=True)
Base = declarative_base()
async_session = sessionmaker(
    engine, class_=AsyncSession, expire_on_commit=False
)

