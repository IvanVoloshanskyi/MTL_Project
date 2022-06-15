from environs import Env


env = Env()
env.read_env()

BOT_TOKEN = env.str("BOT_TOKEN")
ADMINS = env.list("ADMINS")  # Список адмінів (Зробіть БД плз)
IP = env.str("ip")  # Айпі хоста


# Саппорти
support_ids = [

    401057091,
    # 771538156,
]
