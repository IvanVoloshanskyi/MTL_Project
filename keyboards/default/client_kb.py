from aiogram.types import ReplyKeyboardMarkup, ReplyKeyboardRemove, KeyboardButton, InlineKeyboardButton
from aiogram.utils import callback_data

start_button = KeyboardButton('/start', )
about_us_button = KeyboardButton('About us')
support_call_button = KeyboardButton('Support call')
manual_button = KeyboardButton('Manual')
kb_client = ReplyKeyboardMarkup(resize_keyboard=True, one_time_keyboard=False)

# text="About our team"
kb_client.add(about_us_button).add(support_call_button).insert(manual_button)