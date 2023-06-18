## 使用技術

- TypeScript <https://effectivetypescript.com/>
- Svelte(framework) <https://svelte.jp/examples/media-elements>
- Zod(validator) <https://superforms.vercel.app/>
- Svelte Store (state management) <https://qiita.com/sho_U/items/8a508a75aeabed204648>
- Astro(rendering framework) <https://docs.astro.build/ja/getting-started/>
- Storybook(UI catalog) <https://storybook.js.org/>
- Prettier(formatter) <https://docs.astro.build/ja/editor-setup/>
- Eslint(linter) <https://eslint.org/docs/latest/use/getting-started>
- Vitest (test library) <https://vitest.dev/>
- i18next + svelte <https://github.com/NishuGoel/svelte-i18next>
- moment <https://momentjs.com/>

## 環境構築

```
* Makefile 参照
$ make init (compose.override.ymlを適宜書き換えてください)
$ make ui  (docker コンテナ起動)
```

## コーディング規約

- make fmt (PR 前に実行必須)
- make lint (PR 前に実行必須)
- import 規規 (component import 時)
  - atoms -> prefix : "Atom_Components"
  - molecules -> prefix : "Mole_Components"
- 型定義は type を使用 (inteface だと型の拡張ができてしまう)
- コーディング方針 <https://typescript-jp.gitbook.io/deep-dive/styleguide>
- Eslint + Prettier <https://zenn.dev/tetracalibers/articles/5780f9a160addf>
- コーディング規約を変更する場合、".prettierrc.yml", ".eslintrc.yml"を変更する

## Architecture

**_ Atomic Design _**
<https://bradfrost.com/blog/post/atomic-web-design/>

## Directory 構成

```
src
├── api (各APIとのfetchロジック)
├── components (各コンポーネント)
│   ├── atoms (最小単位のコンポーネント)
│   └── molecules (2つ以上のAtomsを組み合わせたコンポーネント)
├── layouts (各ドメインに特化したコンポーネント、ロジックも組み込む)
├── pages ()
├── store (UI状態保持 Svelte Store)
├── stories (storybook)
├── util (env, error, time etc)
└── validate (バリデーションスキーマ　Svelte SuperForm)
```

## test 方針

- storybook test <https://storybook.js.org/tutorials/intro-to-storybook/react/ja/test/>
  - snapshot test
  - regression test
- vitest

```
test配下にテストを作成
srcとtestの構造は連動する
testファイル名: "<テスト対象ファイル名>.spec.ts"
```

## Storybook

作成したコンポーネントを storybook と連動させる
