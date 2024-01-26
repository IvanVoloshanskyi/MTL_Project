from aiogram import types
from aiogram.dispatcher.filters.builtin import CommandStart
from aiogram.types import user

from keyboards.default.client_kb import kb_client
from loader import dp


@dp.message_handler(CommandStart())
async def bot_start(message: types.Message):
    await message.answer(f"Hello {message.from_user.full_name}, it`s MTL support bot.\nYou are welcome",
                         reply_markup=kb_client)
