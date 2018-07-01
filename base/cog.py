from discord.ext import commands

from mydiscord.cogbase import CogBase
from mydiscord.context import Context
from mydiscord.message import Message


class Cog(CogBase):
    @commands.group(invoke_without_command=True)
    async def test(self, ctx: Context) -> None:
        await ctx.acknowledge()

    @test.command()
    async def ping(self, ctx: Context) -> None:
        await ctx.post(Message('pong'))

    @commands.command()
    async def contact(self, ctx: Context) -> None:
        """
        Shows an invite to the bot development server.
        """
        await ctx.post(Message('https://discord.gg/gE9Dh57'))
