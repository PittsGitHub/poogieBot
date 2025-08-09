# 🐽 PoogieBot

A Discord bot for **Monster Hunter Wilds** gear lookups based on desired skill name

### 🔐 Required Permissions

PoogieBot only needs minimal permissions to function properly:

- Send Messages
- Embed Links

These permissions are built into the official invite link:
[Invite PoogieBot](https://discord.com/oauth2/authorize?client_id=1400521755604287582&permissions=83968&integration_type=0&scope=bot)


🧪 Examples:

🛡️ Show all low rank armor for part breaker
`!find low, armor, part breaker`

📿 Show talismans with part breaker
`!find , talisman, part breaker`

💎 Show decorations with critical boost
`!find , decoration, critical boost`

🗡️ Show rarity 8 greatswords with critical draw
`!find 8, greatsword, critical draw`

---

## ✨ Features

- 🔍 Skill-based lookup for:
  - Armor Sets (with set/group bonuses and piece breakdowns)
  - Decorations
  - Talismans
  - Weapons (with stats)
- 📎 Automatic wiki links for every item
- 🧽 Input cleansing for ease of use

---

## 📜 Command Parameters

All commands are prefixed with `!find` then parameters must be in the correct order and comma seperated
`!find [rarity], [type], [skill name]`

👑 Rarity

    high — Matches rarity 5–8 (default)

    low — Matches rarity 0–4

    A specific rarity number like 7

    Ranges:

        "5+" — Rarity 5 and up

        "4-" — Rarity 4 and down

🔨 Type

    A category:

        armor, decoration, talisman, weapons

    Or a specific weapon type:

        greatsword, longsword, bow, etc.

🪤 Skill Name

    critical eye

    part breaker

    windproof

    evade extender
