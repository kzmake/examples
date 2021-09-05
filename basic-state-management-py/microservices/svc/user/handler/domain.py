from dataclasses import dataclass

from dataclasses_json import dataclass_json


# 値オブジェクト
@dataclass_json
@dataclass(frozen=True)
class UserID:
    value: str


@dataclass_json
@dataclass(frozen=True)
class UserName:
    value: str


# エンティティ
@dataclass_json
@dataclass
class User:
    user_id: UserID

    name: UserName

    def rename(self, name: UserName) -> None:
        self.name = name


# 集約
AggregateRootUser = User
