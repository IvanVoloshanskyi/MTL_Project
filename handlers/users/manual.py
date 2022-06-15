from aiogram import types
from aiogram.dispatcher.filters.builtin import CommandStart
from aiogram.dispatcher.filters import Text
from keyboards.default.client_kb import kb_client
from loader import dp, bot


@dp.message_handler(Text(equals="Manual"))
async def manual(message: types.Message):
    # doc = open('Manual_MTL' + '.docx')
    await message.answer_document(open('files/Manual_MTL.docx', 'rb'))
