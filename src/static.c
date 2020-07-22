#include "static.h"
#include "common/object/modifier.h"

#include "client/drawable/drawdb.h"

#include "proto.h"

void nullsub_38() {}
void nullsub_39() {}
void nullsub_40() {}
void nullsub_41() {}
void nullsub_42() {}
void nullsub_43() {}
void nullsub_44() {}

const char* table_5184[] = {"NULL",
							"AnchorCast",
							"AnchorOn",
							"AnchorOff",
							"BlindCast",
							"BlindOn",
							"BlindOff",
							"BlinkCast",
							"BurnCast",
							"CleansingFlameCast",
							"ChannelLifeCast",
							"ChannelLifeEffect",
							"ChannelLifeStop",
							"CancelCast",
							"CharmCast",
							"CharmSuccess",
							"CharmFailure",
							"ConfuseCast",
							"ConfuseOn",
							"ConfuseOff",
							"CounterspellCast",
							"CurePoisonCast",
							"CurePoisonEffect",
							"DeathCast",
							"DeathOn",
							"DeathOff",
							"DeathTimer",
							"DeathRayCast",
							"DetonateGlyphCast",
							"DrainManaCast",
							"EarthquakeCast",
							"EnergyBoltCast",
							"EnergyBoltSustain",
							"FearCast",
							"FearOn",
							"FearOff",
							"ForceOfNatureCast",
							"ForceOfNatureReflect",
							"ForceOfNatureRelease",
							"GlyphCast",
							"GlyphDetonate",
							"FireballCast",
							"FireballExplode",
							"FirewalkCast",
							"FirewalkOn",
							"FirewalkOff",
							"FirewalkFlame",
							"FistCast",
							"FistHit",
							"ForceFieldCast",
							"FrostCast",
							"FrostOn",
							"FrostOff",
							"FumbleCast",
							"FumbleEffect",
							"GreaterHealCast",
							"GreaterHealEffect",
							"GreaterHealStop",
							"HasteCast",
							"HasteOn",
							"HasteOff",
							"InfravisionCast",
							"InfravisionOn",
							"InfraVisionOff",
							"InversionCast",
							"InvisibilityCast",
							"InvisibilityOn",
							"InvisibilityOff",
							"InvulnerabilityCast",
							"InvulnerabilityOn",
							"InvulnerabilityOff",
							"InvulnerableEffect",
							"LesserHealCast",
							"LesserHealEffect",
							"LightCast",
							"LightOn",
							"LightOff",
							"LightningCast",
							"LightningBolt",
							"LockCast",
							"ManaBombCast",
							"ManaBombEffect",
							"MarkCast",
							"MagicMissileCast",
							"MagicMissileDetonate",
							"MeteorCast",
							"MeteorShowerCast",
							"MeteorHit",
							"MoonglowCast",
							"MoonglowOn",
							"MoonglowOff",
							"NullifyCast",
							"NullifyOn",
							"NullifyOff",
							"PhantomCast",
							"PixieSwarmCast",
							"PixieHit",
							"PlasmaCast",
							"PlasmaSustain",
							"PoisonCast",
							"PoisonEffect",
							"ProtectionFromFireCast",
							"ProtectionFromFireOn",
							"ProtectionFromFireOff",
							"ProtectionFromFireEffect",
							"ProtectionFromElectricityCast",
							"ProtectionFromElectricityOn",
							"ProtectionFromElectricityOff",
							"ProtectionFromElectricityEffect",
							"ProtectionFromPoisonCast",
							"ProtectionFromPoisonOn",
							"ProtectionFromPoisonOff",
							"ProtectionFromPoisonEffect",
							"ProtectionFromMagicCast",
							"ProtectionFromMagicOn",
							"ProtectionFromMagicOff",
							"ProtectionFromMagicEffect",
							"PullCast",
							"PushCast",
							"ReflectiveShieldCast",
							"ReflectiveShieldOn",
							"ReflectiveShieldOff",
							"ReflectiveShieldEffect",
							"RegenerationOn",
							"RegenerationOff",
							"RunCast",
							"RunOn",
							"RunOff",
							"ShieldCast",
							"ShieldOn",
							"ShieldOff",
							"ShieldRepelled",
							"ShockCast",
							"ShockOn",
							"ShockOff",
							"Shocked",
							"SlowCast",
							"SlowOn",
							"SlowOff",
							"StunCast",
							"StunOn",
							"StunOff",
							"SummonCast",
							"SwapCast",
							"TagCast",
							"TagOn",
							"TagOff",
							"TeleportOut",
							"TeleportIn",
							"TeleportToMarkerCast",
							"TeleportToMarkerEffect",
							"TeleportPopCast",
							"TeleportToTargetCast",
							"TelekinesisCast",
							"TelekinesisOn",
							"TelekinesisOff",
							"ToxicCloudCast",
							"TriggerGlyphCast",
							"TurnUndeadCast",
							"TurnUndeadEffect",
							"VampirismCast",
							"VampirismOn",
							"VampirismOff",
							"DrainHealth",
							"VillainCast",
							"VillainOn",
							"VillainOff",
							"WallCast",
							"WallOn",
							"WallOff",
							"BerserkerChargeInvoke",
							"BerserkerCrash",
							"BerserkerChargeOn",
							"BerserkerChargeOff",
							"WarcryInvoke",
							"WarcryOn",
							"WarcryOff",
							"HarpoonInvoke",
							"HarpoonOn",
							"HarpoonOff",
							"TreadLightlyInvoke",
							"TreadLightlyOn",
							"TreadLightlyOff",
							"EyeOfTheWolfInvoke",
							"EyeOfTheWolfOn",
							"EyeOfTheWolfOff",
							"SpellPhonemeUp",
							"SpellPhonemeUpRight",
							"SpellPhonemeRight",
							"SpellPhonemeDownRight",
							"SpellPhonemeDown",
							"SpellPhonemeDownLeft",
							"SpellPhonemeLeft",
							"SpellPhonemeUpLeft",
							"FemaleSpellPhonemeUp",
							"FemaleSpellPhonemeUpRight",
							"FemaleSpellPhonemeRight",
							"FemaleSpellPhonemeDownRight",
							"FemaleSpellPhonemeDown",
							"FemaleSpellPhonemeDownLeft",
							"FemaleSpellPhonemeLeft",
							"FemaleSpellPhonemeUpLeft",
							"NPCSpellPhonemeUp",
							"NPCSpellPhonemeUpRight",
							"NPCSpellPhonemeRight",
							"NPCSpellPhonemeDownRight",
							"NPCSpellPhonemeDown",
							"NPCSpellPhonemeDownLeft",
							"NPCSpellPhonemeLeft",
							"NPCSpellPhonemeUpLeft",
							"NPCFemaleSpellPhonemeUp",
							"NPCFemaleSpellPhonemeUpRight",
							"NPCFemaleSpellPhonemeRight",
							"NPCFemaleSpellPhonemeDownRight",
							"NPCFemaleSpellPhonemeDown",
							"NPCFemaleSpellPhonemeDownLeft",
							"NPCFemaleSpellPhonemeLeft",
							"NPCFemaleSpellPhonemeUpLeft",
							"FireballWand",
							"SmallFireballWand",
							"FlareWand",
							"LightningWand",
							"DepletedWand",
							"Ricochet",
							"WeaponEffectFire",
							"WeaponEffectElectricity",
							"AwardSpell",
							"AwardGuide",
							"BlueFireDamage",
							"DrainMana",
							"ManaRecharge",
							"PermanentFizzle",
							"ManaEmpty",
							"Lock",
							"Unlock",
							"ElevEnable",
							"ElevDisable",
							"OpenWoodenDoor",
							"MoveWoodenDoor",
							"CloseWoodenDoor",
							"WoodenDoorLocked",
							"OpenGate",
							"MoveGate",
							"CloseGate",
							"GateLocked",
							"OpenWoodenGate",
							"CloseWoodenGate",
							"OpenHeavyWoodenDoor",
							"CloseHeavyWoodenDoor",
							"ElevStoneUp",
							"ElevStoneDown",
							"ElevWoodUp",
							"ElevWoodDown",
							"ElevMechUp",
							"ElevMechDown",
							"ElevLavaUp",
							"ElevLavaDown",
							"ElevGreenUp",
							"ElevGreenDown",
							"ElevLOTDUp",
							"ElevLOTDDown",
							"Gear1",
							"Gear2",
							"Gear3",
							"SmallRockMove",
							"MediumRockMove",
							"LargeRockMove",
							"BoulderMove",
							"WalkOnSnow",
							"WalkOnStone",
							"WalkOnDirt",
							"WalkOnWood",
							"WalkOnWater",
							"WalkOnMud",
							"RunOnSnow",
							"RunOnStone",
							"RunOnDirt",
							"RunOnWood",
							"RunOnWater",
							"RunOnMud",
							"PlayerFallThud",
							"BarrelMove",
							"BlackPowderBurn",
							"FireExtinguish",
							"PolypExplode",
							"PowderBarrelExplode",
							"BarrelBreak",
							"WaterBarrelBreak",
							"LOTDBarrelBreak",
							"WineCaskBreak",
							"BarrelStackBreak",
							"CoffinBreak",
							"WaspHiveBreak",
							"WorkstationBreak",
							"CrushLight",
							"CrushMedium",
							"CrushHard",
							"SentryRayHitWall",
							"SentryRayHit",
							"DeathRayKill",
							"TauntLaugh",
							"TauntShakeFist",
							"TauntPoint",
							"FlagPickup",
							"FlagDrop",
							"FlagRespawn",
							"FlagCapture",
							"TreasurePickup",
							"TreasureDrop",
							"GameOver",
							"ServerOptionsChange",
							"CamperAlarm",
							"PlayerEliminated",
							"CrownChange",
							"HumanMaleEatFood",
							"HumanMaleEatApple",
							"HumanMaleDrinkJug",
							"HumanMaleHurtLight",
							"HumanMaleHurtMedium",
							"HumanMaleHurtHeavy",
							"HumanMaleHurtPoison",
							"HumanMaleDie",
							"HumanMaleExertionLight",
							"HumanMaleExertionHeavy",
							"HumanFemaleEatFood",
							"HumanFemaleEatApple",
							"HumanFemaleDrinkJug",
							"HumanFemaleHurtLight",
							"HumanFemaleHurtMedium",
							"HumanFemaleHurtHeavy",
							"HumanFemaleHurtPoison",
							"HumanFemaleDie",
							"HumanFemaleExertionLight",
							"HumanFemaleExertionHeavy",
							"MonsterEatFood",
							"BatMove",
							"BatBite",
							"BatDie",
							"BatRecognize",
							"SmallSpiderMove",
							"SmallSpiderIdle",
							"SmallSpiderBite",
							"SmallSpiderDie",
							"SmallSpiderRecognize",
							"LargeSpiderMove",
							"LargeSpiderIdle",
							"LargeSpiderBite",
							"LargeSpiderDie",
							"LargeSpiderSpit",
							"LargeSpiderRecognize",
							"LargeSpiderHurt",
							"WebGooHit",
							"SkeletonMove",
							"SkeletonAttackInit",
							"SkeletonHitting",
							"SkeletonMissing",
							"SkeletonDie",
							"SkeletonRecognize",
							"SkeletonHurt",
							"SkeletonLordMove",
							"SkeletonLordAttackInit",
							"SkeletonLordHitting",
							"SkeletonLordMissing",
							"SkeletonLordDie",
							"SkeletonLordRecognize",
							"SkeletonLordHurt",
							"BomberRecognize",
							"BomberSummon",
							"BomberDie",
							"BomberMove",
							"DemonRecognize",
							"DemonTaunt",
							"DemonSpellInit",
							"DemonDie",
							"DemonMove",
							"DemonHurt",
							"GruntMove",
							"GruntTaunt",
							"GruntIdle",
							"GruntAttackInit",
							"GruntHitting",
							"GruntMissing",
							"GruntDie",
							"GruntBowTwang",
							"GruntRecognize",
							"GruntHurt",
							"OgreBruteListen",
							"OgreBruteTaunt",
							"OgreBruteIdle",
							"OgreBruteEngage",
							"OgreBruteRecognize",
							"OgreBruteMove",
							"OgreBruteHurt",
							"OgreBruteDie",
							"OgreBruteMeleeHit",
							"OgreBruteMeleeMiss",
							"OgreBruteAttackInit",
							"OgreWarlordListen",
							"OgreWarlordTaunt",
							"OgreWarlordIdle",
							"OgreWarlordEngage",
							"OgreWarlordRecognize",
							"OgreWarlordMove",
							"OgreWarlordHurt",
							"OgreWarlordDie",
							"OgreWarlordMeleeHit",
							"OgreWarlordMeleeMiss",
							"OgreWarlordAttackInit",
							"OgreWarlordThrow",
							"ScorpionRecognize",
							"ScorpionMove",
							"ScorpionIdle",
							"ScorpionAttackInit",
							"ScorpionStingHit",
							"ScorpionStingMiss",
							"ScorpionClawHit",
							"ScorpionClawMiss",
							"ScorpionHurt",
							"ScorpionDie",
							"LeechMove",
							"LeechIdle",
							"LeechAttackInit",
							"LeechHitting",
							"LeechMissing",
							"LeechDie",
							"LeechRecognize",
							"LeechHurt",
							"BearMove",
							"BearAttackInit",
							"BearHitting",
							"BearMissing",
							"BearDie",
							"BearRecognize",
							"BearHurt",
							"WolfMove",
							"WolfIdle",
							"WolfAttackInit",
							"WolfHitting",
							"WolfMissing",
							"WolfDie",
							"WolfRecognize",
							"WolfHurt",
							"WolfHowl",
							"WolfGrowl",
							"PlantSleep",
							"PlantAttackInit",
							"PlantHitting",
							"PlantMissing",
							"PlantDie",
							"PlantRecognize",
							"PlantHurt",
							"PlantGrowl",
							"MimicMove",
							"MimicIdle",
							"MimicAttackInit",
							"MimicHitting",
							"MimicMissing",
							"MimicDie",
							"MimicRecognize",
							"MimicHurt",
							"MimicMorph",
							"ZombieMove",
							"ZombieIdle",
							"ZombieAttackInit",
							"ZombieHitting",
							"ZombieMissing",
							"ZombieDie",
							"ZombieRecognize",
							"ZombieHurt",
							"ZombieGetUp",
							"VileZombieMove",
							"VileZombieIdle",
							"VileZombieAttackInit",
							"VileZombieHitting",
							"VileZombieMissing",
							"VileZombieDie",
							"VileZombieRecognize",
							"VileZombieHurt",
							"VileZombieGetUp",
							"ImpMove",
							"ImpShoot",
							"ImpSteal",
							"ImpDie",
							"ImpRecognize",
							"GolemMove",
							"GolemHitting",
							"GolemMissing",
							"GolemDie",
							"GolemHurt",
							"GolemRecognize",
							"MechGolemMove",
							"MechGolemAttackInit",
							"MechGolemHitting",
							"MechGolemMissing",
							"MechGolemDie",
							"MechGolemHurt",
							"MechGolemRecognize",
							"TowerRecognize",
							"TowerShoot",
							"SkullRecognize",
							"SkullShoot",
							"GhostMove",
							"GhostHitting",
							"GhostDie",
							"GhostHurt",
							"GhostRecognize",
							"WizardTalkable",
							"WizardMove",
							"WizardDie",
							"WizardHurt",
							"WizardRecognize",
							"WizardEngage",
							"WizardRetreat",
							"WaspMove",
							"WaspIdle",
							"WaspSting",
							"WaspDie",
							"WaspRecognize",
							"EmberDemonMove",
							"EmberDemonTaunt",
							"EmberDemonHitting",
							"EmberDemonMissing",
							"EmberDemonDie",
							"EmberDemonThrow",
							"EmberDemonRecognize",
							"EmberDemonHurt",
							"UrchinMove",
							"UrchinTaunt",
							"UrchinIdle",
							"UrchinDie",
							"UrchinThrow",
							"UrchinRecognize",
							"UrchinHurt",
							"UrchinFlee",
							"UrchinShamanMove",
							"UrchinShamanTaunt",
							"UrchinShamanIdle",
							"UrchinShamanDie",
							"UrchinShamanRecognize",
							"UrchinShamanHurt",
							"UrchinShamanFlee",
							"ArcherMove",
							"ArcherTaunt",
							"ArcherIdle",
							"ArcherDie",
							"ArcherMissileInit",
							"ArcherShoot",
							"ArcherRecognize",
							"ArcherHurt",
							"ArcherRetreat",
							"SwordsmanMove",
							"SwordsmanTaunt",
							"SwordsmanIdle",
							"SwordsmanAttackInit",
							"SwordsmanHitting",
							"SwordsmanMissing",
							"SwordsmanDie",
							"SwordsmanRecognize",
							"SwordsmanHurt",
							"SwordsmanRetreat",
							"BeholderMove",
							"BeholderIdle",
							"BeholderAttackInit",
							"BeholderDie",
							"BeholderRecognize",
							"BeholderHurt",
							"DryadMove",
							"DryadTaunt",
							"DryadDie",
							"DryadRecognize",
							"DryadHurt",
							"EvilCherubMove",
							"EvilCherubTaunt",
							"EvilCherubIdle",
							"EvilCherubMissileInit",
							"EvilCherubShoot",
							"EvilCherubDie",
							"EvilCherubRecognize",
							"EvilCherubHurt",
							"FishDie",
							"FrogDie",
							"FrogRecognize",
							"HecubahTaunt",
							"HecubahTalkable",
							"HecubahMove",
							"HecubahAttackInit",
							"HecubahDie",
							"HecubahRecognize",
							"HecubahHurt",
							"HecubahDieFrame0A",
							"HecubahDieFrame0B",
							"HecubahDieFrame98",
							"HecubahDieFrame194",
							"HecubahDieFrame283",
							"HecubahDieFrame439",
							"NecromancerTaunt",
							"NecromancerTalkable",
							"NecromancerMove",
							"NecromancerAttackInit",
							"NecromancerDie",
							"NecromancerRecognize",
							"NecromancerEngage",
							"NecromancerRetreat",
							"NecromancerHurt",
							"LichMove",
							"LichAttackInit",
							"LichDie",
							"LichRecognize",
							"LichHurt",
							"FlyingGolemMove",
							"FlyingGolemShoot",
							"FlyingGolemDie",
							"FlyingGolemRecognize",
							"FlyingGolemHurt",
							"NPCTalkable",
							"NPCIdle",
							"NPCDie",
							"NPCRecognize",
							"NPCRetreat",
							"NPCHurt",
							"MaidenIdle",
							"MaidenTalkable",
							"MaidenFlee",
							"MaidenDie",
							"MaidenHurt",
							"RatDie",
							"ShadeMove",
							"ShadeAttackInit",
							"ShadeDie",
							"ShadeRecognize",
							"ShadeHurt",
							"WeirdlingMove",
							"WillOWispMove",
							"WillOWispIdle",
							"WillOWispDie",
							"WillOWispHurt",
							"WillOWispRecognize",
							"WillOWispEngage",
							"TrollMove",
							"TrollIdle",
							"TrollDie",
							"TrollHurt",
							"TrollRecognize",
							"TrollAttackInit",
							"TrollFlatus",
							"MaleNPC1Talkable",
							"MaleNPC1Idle",
							"MaleNPC1Recognize",
							"MaleNPC1Engage",
							"MaleNPC1Retreat",
							"MaleNPC1Hurt",
							"MaleNPC1Die",
							"MaleNPC1AttackInit",
							"MaleNPC2Talkable",
							"MaleNPC2Idle",
							"MaleNPC2Recognize",
							"MaleNPC2Engage",
							"MaleNPC2Retreat",
							"MaleNPC2Hurt",
							"MaleNPC2Die",
							"MaleNPC2AttackInit",
							"Maiden1Talkable",
							"Maiden1Idle",
							"Maiden1Recognize",
							"Maiden1Retreat",
							"Maiden1Hurt",
							"Maiden1Die",
							"Maiden1AttackInit",
							"Maiden2Talkable",
							"Maiden2Idle",
							"Maiden2Recognize",
							"Maiden2Retreat",
							"Maiden2Hurt",
							"Maiden2Die",
							"Maiden2AttackInit",
							"HorvathTalkable",
							"HorvathEngage",
							"HorvathHurt",
							"HorvathDie",
							"Wizard1Talkable",
							"Wizard1Idle",
							"Wizard1Recognize",
							"Wizard1Engage",
							"Wizard1Retreat",
							"Wizard1Hurt",
							"Wizard1Die",
							"Wizard1AttackInit",
							"Wizard2Talkable",
							"Wizard2Idle",
							"Wizard2Recognize",
							"Wizard2Engage",
							"Wizard2Retreat",
							"Wizard2Hurt",
							"Wizard2Die",
							"Wizard2AttackInit",
							"FireKnight1Talkable",
							"FireKnight1Idle",
							"FireKnight1Recognize",
							"FireKnight1Engage",
							"FireKnight1Retreat",
							"FireKnight1Hurt",
							"FireKnight1Die",
							"FireKnight1AttackInit",
							"FireKnight2Talkable",
							"FireKnight2Idle",
							"FireKnight2Recognize",
							"FireKnight2Engage",
							"FireKnight2Retreat",
							"FireKnight2Hurt",
							"FireKnight2Die",
							"FireKnight2AttackInit",
							"Guard1Talkable",
							"Guard1Idle",
							"Guard1Recognize",
							"Guard1Engage",
							"Guard1Retreat",
							"Guard1Hurt",
							"Guard1Die",
							"Guard1AttackInit",
							"Guard2Talkable",
							"Guard2Idle",
							"Guard2Recognize",
							"Guard2Engage",
							"Guard2Retreat",
							"Guard2Hurt",
							"Guard2Die",
							"Guard2AttackInit",
							"WoundedNPCTalkable",
							"WoundedNPCIdle",
							"WoundedNPCRecognize",
							"WoundedNPCRetreat",
							"WoundedNPCHurt",
							"WoundedNPCDie",
							"WoundedNPCAttackInit",
							"HorrendousTalkable",
							"HorrendousRecognize",
							"HorrendousRetreat",
							"HorrendousHurt",
							"HorrendousDie",
							"HorrendousAttackInit",
							"SecretWallOpen",
							"SecretWallClose",
							"SecretWallEarthOpen",
							"SecretWallEarthClose",
							"SecretWallMetalOpen",
							"SecretWallMetalClose",
							"SecretWallStoneOpen",
							"SecretWallStoneClose",
							"SecretWallWoodOpen",
							"SecretWallWoodClose",
							"TriggerPressed",
							"TriggerReleased",
							"PotionUse",
							"PotionBreak",
							"RestoreHealth",
							"RestoreMana",
							"WallDestroyed",
							"WallDestroyedStone",
							"WallDestroyedWood",
							"WallDestroyedMetal",
							"ChestOpen",
							"CryptChestOpen",
							"SackChestOpen",
							"ChestLocked",
							"EggBreak",
							"Poisoned",
							"ButtonPress",
							"ButtonRelease",
							"LeverToggle",
							"SwitchToggle",
							"ChainPull",
							"ShortBellsUp",
							"LongBellsUp",
							"LongBellsDown",
							"BigBell",
							"MetallicBong",
							"Chime",
							"BigGong",
							"SmallGong",
							"MysticChant",
							"TripleChime",
							"Clank1",
							"Clank2",
							"Clank3",
							"MapOpen",
							"MapClose",
							"BookOpen",
							"BookClose",
							"PageTurn",
							"InventoryOpen",
							"InventoryClose",
							"InventoryPickup",
							"InventoryDrop",
							"SpellPickup",
							"SpellDrop",
							"SpellPopOffBook",
							"TrapEditorOpen",
							"TrapEditorClose",
							"ChangeSpellbar",
							"ExpandSpellbar",
							"CollapseSpellbar",
							"CreatureCageAppears",
							"CreatureCageHides",
							"ShopRepairItem",
							"MetalArmorPickup",
							"MetalArmorDrop",
							"MetalArmorBreak",
							"LeatherArmorPickup",
							"LeatherArmorDrop",
							"LeatherArmorBreak",
							"WoodenArmorPickup",
							"WoodenArmorDrop",
							"WoodenArmorBreak",
							"ClothArmorPickup",
							"ClothArmorDrop",
							"ClothArmorBreak",
							"ShoesPickup",
							"ShoesDrop",
							"MetalWeaponBreak",
							"WoodWeaponBreak",
							"KeyPickup",
							"KeyDrop",
							"AmuletPickup",
							"AmuletDrop",
							"TrapPickup",
							"TrapDrop",
							"BookPickup",
							"BookDrop",
							"ScrollPickup",
							"ScrollDrop",
							"WandPickup",
							"WandDrop",
							"PotionPickup",
							"PotionDrop",
							"MeatPickup",
							"MeatDrop",
							"ApplePickup",
							"AppleDrop",
							"ShroomPickup",
							"ShroomDrop",
							"SpectaclesPickup",
							"SpectaclesDrop",
							"MetalWeaponPickup",
							"MetalWeaponDrop",
							"WoodenWeaponPickup",
							"WoodenWeaponDrop",
							"BearTrapTriggered",
							"PoisonTrapTriggered",
							"StoneHitStone",
							"StoneHitEarth",
							"StoneHitWood",
							"StoneHitMetal",
							"StoneHitFlesh",
							"WoodHitStone",
							"WoodHitEarth",
							"WoodHitWood",
							"WoodHitMetal",
							"WoodHitFlesh",
							"MetalHitStone",
							"MetalHitEarth",
							"MetalHitWood",
							"MetalHitMetal",
							"MetalHitFlesh",
							"FleshHitStone",
							"FleshHitEarth",
							"FleshHitWood",
							"FleshHitMetal",
							"FleshHitFlesh",
							"DiamondHitStone",
							"DiamondHitEarth",
							"DiamondHitWood",
							"DiamondHitMetal",
							"DiamondHitFlesh",
							"HitStoneBreakable",
							"HitEarthBreakable",
							"HitWoodBreakable",
							"HitMetalBreakable",
							"HitMagicBreakable",
							"HitMetalShield",
							"PunchMissing",
							"LongswordMissing",
							"SwordMissing",
							"HammerMissing",
							"AxeMissing",
							"MaceMissing",
							"BowShoot",
							"CrossBowShoot",
							"BowEmpty",
							"CrossBowEmpty",
							"ArrowTrapShoot",
							"GreatSwordReflect",
							"ChakramThrow",
							"ChakramCatch",
							"ChakramFallToGround",
							"StaffBlock",
							"NextWeapon",
							"HeartBeat",
							"GenerateTick",
							"SummonClick",
							"SummonComplete",
							"SummonAbort",
							"ManaClick",
							"LevelUp",
							"JournalEntryAdd",
							"SecretFound",
							"EarthRumbleMajor",
							"EarthRumbleMinor",
							"ElectricalArc1",
							"FloorSpikesUp",
							"FloorSpikesDown",
							"SpikeBlockMove",
							"BoulderRoll",
							"ArcheryContestBegins",
							"HorrendousIsKilled",
							"StaffOblivionAchieve1",
							"StaffOblivionAchieve2",
							"StaffOblivionAchieve3",
							"StaffOblivionAchieve4",
							"FireGrate",
							"MechGolemPowerUp",
							"ShellSelect",
							"ShellClick",
							"ShellSlideIn",
							"ShellSlideOut",
							"ShellMouseBoom",
							"NoCanDo",
							"BallThrow",
							"BallGrab",
							"BallBounce",
							"BallHitGoal",
							"BirdAmbient1",
							"LoopAmbienceCastle",
							"LoopAmbienceCave",
							"LoopAmbienceForest",
							"LoopAmbienceInferno",
							"LoopAmbienceLOTD",
							"LoopAmbienceMine",
							"LoopAmbienceSwamp",
							"LoopAmbBird2",
							"LoopAmbBirdSwamp1",
							"LoopAmbBirdSwamp2",
							"LoopAmbBrook2",
							"LoopAmbBrook3",
							"LoopAmbCricket1",
							"LoopAmbCricket2",
							"LoopAmbCricket3",
							"LoopAmbCrow1",
							"LoopAmbCrow2",
							"LoopAmbCrowWL",
							"LoopAmbDog1",
							"LoopAmbDripCave1",
							"LoopAmbDripCave2",
							"LoopAmbDripRoom",
							"LoopAmbFrogSwamp1",
							"LoopAmbFrogSwamp2",
							"LoopAmbFrogSwamp3",
							"LoopAmbFrogSwamp4",
							"LoopAmbFrogSwamp5",
							"LoopAmbFliesSwamp",
							"LoopAmbRodentForest1",
							"LoopAmbRodentForest2",
							"LoopAmbCaveRumble",
							"LoopAmbSpiritsFOV",
							"LoopAmbSpiritsLOTD",
							"LoopAmbWindCave1",
							"LoopAmbWindCave2",
							"LoopAmbWindCave3",
							"LoopAmbWindLOTD",
							"LoopAmbWindWL1",
							"LoopAmbWindWL2",
							"LoopAmbHeartNox",
							"LoopAmbBird3",
							"LoopAmbBird4",
							"LoopAmbBird5",
							"LoopAmbBird6",
							"LoopAmbBird7",
							"LoopAmbBird8",
							"LoopAmbIxTemple",
							"LoopAmbLeaves",
							"LoopAmbHowls",
							"LoopAmbWeirdling",
							"LoopAmbSpikeWave",
							"LoopAmbBeachBirds",
							"LoopAmbBeachWaves",
							"LoopCampfire",
							"LoopTorch",
							"LoopBlueFire",
							"LoopWorkstation",
							"LoopWaterFountain",
							"LoopWindmill",
							"LoopOrrery",
							"LoopVandegraf",
							"LoopGearSmall",
							"LoopGearLarge",
							"LoopGearPulley",
							"LoopPolyp",
							"LoopSentryGlobe",
							"LoopSpikePillar",
							"HarpoonBroken",
							"HarpoonReel",
							"MonsterGeneratorDie",
							"MonsterGeneratorHurt",
							"MonsterGeneratorSpawn",
							"PlayerExit",
							"AwardLife",
							"SoulGateTouch",
							"QuestRespawn",
							"QuestFinalDeath",
							"QuestPlayerJoinGame",
							"QuestStatScreen",
							"QuestIntroScreen",
							"QuestPlayerExitGame",
							"QuestLockedChest",
							"LoopMonsterGenerator",
							"StoneDoorOpen",
							"StoneDoorClose",
							"DiamondDrop",
							"DiamondPickup",
							"RubyDrop",
							"RubyPickup",
							"EmeraldDrop",
							"EmeraldPickup",
							"EnableSharedKeyMode"};

table_26792_t table_26792[] = {
	{"DamageMultiplierEffect", &sub_4E04C0, &sub_4235C0},
	{"StunEffect", &sub_4E04D0, &sub_423780},
	{"FireEffect", &sub_4E0550, &sub_4235C0},
	{"FireRingEffect", &sub_4E05B0, &sub_423780},
	{"BlueFireRingEffect", &sub_4E05F0, &sub_423780},
	{"FrostEffect", &nullsub_38, &sub_423780},
	{"RecoilEffect", &sub_4E0640, &sub_4235C0},
	{"ConfuseEffect", &sub_4E0670, &sub_423780},
	{"LightningEffect", &sub_4E06F0, &sub_4235C0},
	{"DrainManaEffect", &sub_4E0740, &sub_4235C0},
	{"VampirismEffect", &sub_4E07C0, &sub_4235C0},
	{"PoisonEffect", &sub_4E0850, &sub_423780},
	{"PanicEffect", &nullsub_39, &sub_423780},
	{"SympathyEffect", &sub_4E08E0, &sub_4235C0},
	{"ReadinessEffect", &nullsub_22, &sub_423780},
	{"ProjectileSpeedEffect", &sub_4E09B0, &sub_4235C0},
	{"ReplenishmentEffect", &nullsub_36, &sub_423780},
};
int table_26792_cnt = sizeof(table_26792) / sizeof(table_26792_t);

table_27008_t table_27008[] = {
	{"ArmorMultiplierEffect", &sub_4E0370, &sub_4235C0},
	{"DurabilityMultiplierEffect", &sub_4E0380, &sub_4235C0},
	{"ResilienceEffect", &nullsub_40, &sub_4235C0},
	{"InversionEffect", &sub_4E03D0, &sub_423780},
	{"GripEffect", &sub_4E0480, &sub_423780},
	{"BreakingEffect", &nullsub_41, &sub_4235C0},
	{"PunctureProneEffect", &nullsub_42, &sub_4235C0},
};
int table_27008_cnt = sizeof(table_27008) / sizeof(table_27008_t);

table_27104_t table_27104[] = {
	{"RegenerationUpdate", &sub_4E01D0, &sub_423780},
	{"ParasiteUpdate", &nullsub_43, &sub_423780},
	{"AttractionUpdate", &nullsub_44, &sub_423780},
	{"ContinualReplenishmentUpdate", &sub_4E02C0, &sub_423780},
};
int table_27104_cnt = sizeof(table_27104) / sizeof(table_27104_t);

table_27168_t table_27168[] = {
	{"BrillianceEngage", &sub_4DFB50, &sub_423780},
	{"BrillianceDisengage", &sub_4DFB80, &sub_423780},
	{"SpeedEngage", &sub_4DFC30, &sub_4235C0},
	{"SpeedDisengage", &sub_4DFCA0, &sub_4235C0},
	{"FireProtectEngage", &sub_4DFD10, &sub_4235C0},
	{"FireProtectDisengage", &sub_4DFD40, &sub_4235C0},
	{"LightningProtectEngage", &sub_4DFD80, &sub_4235C0},
	{"LightningProtectDisengage", &sub_4DFDB0, &sub_4235C0},
	{"PoisonProtectEngage", &sub_4DFDE0, &sub_4235C0},
	{"PoisonProtectDisengage", &sub_4DFE10, &sub_4235C0},
	{"RegenerationEngage", &sub_4E0140, 0},
	{"RegenerationDisengage", &sub_4E0170, 0},
};
int table_27168_cnt = sizeof(table_27168) / sizeof(table_27168_t);

table_28760_t table_28760[] = {
	{"DESC", &sub_412100_parse_desc},
	{"PRIMARYDESC", &sub_412100_parse_desc},
	{"SECONDARYDESC", &sub_4121B0_parse_second_desc},
	{"IDENTIFYDESC", &sub_412260_parse_ident_desc},
	{"WORTH", &sub_412310_parse_worth},
	{"COLOR", &sub_412360_parse_color},
	{"ATTACKEFFECT", &sub_412380_parse_attack_effect},
	{"ATTACKPREHITEFFECT", &sub_412380_parse_attack_effect},
	{"ATTACKPREDAMAGEEFFECT", &sub_412380_parse_attack_effect},
	{"DEFENDEFFECT", &sub_412490_parse_defend_effect},
	{"DEFENDCOLLIDEEFFECT", &sub_412490_parse_defend_effect},
	{"ENGAGEEFFECT", &sub_412580_parse_engage_effect},
	{"DISENGAGEEFFECT", &sub_412580_parse_engage_effect},
	{"UPDATEEFFECT", &sub_412670_parse_update_effect},
	{"ALLOWED_WEAPONS", &sub_412740_parse_allowed_weapons},
	{"ALLOWED_ARMOR", &sub_4128A0_parse_allowed_armor},
	{"ALLOWED_POSITION", &sub_4128C0_parse_allowed_pos},
};
int table_28760_cnt = sizeof(table_28760) / sizeof(table_28760_t);

nox_video_mode nox_video_modes[] = {
	{640, 480, 0},
	{800, 600, 1},
	{1024, 768, 2},
};
int nox_video_modes_cnt = sizeof(nox_video_modes) / sizeof(nox_video_mode);

#define NOX_VERSION_STRING "1.2b"
char nox_version_string_180[7] = NOX_VERSION_STRING;
char nox_version_string_102944[7] = NOX_VERSION_STRING;
