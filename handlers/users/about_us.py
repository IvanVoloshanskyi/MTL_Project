from aiogram import types
from aiogram.dispatcher.filters.builtin import CommandHelp, Text
from keyboards.default.client_kb import kb_client
from loader import dp


@dp.message_handler(Text(equals="About us"))
async def bot_about_us(message: types.Message):
    text = (f'Hello, we are MTL. Our team consists of 5 people.\n'
                                                 f'Maryan Petlovaniy - AI Dev \n'
                                                 f'Taras Tovarnitskiy - Designer \ Site Dev\n'
                                                 f'Taras Florko - Hardware Dev \n'
                                                 f'Alexander Minaev - Mobile app Dev \n'
                                                 f'Ivan Voloshanskyi - Python bot Developer',
                           )

    await message.answer("\n".join(text), reply_markup=kb_client)
