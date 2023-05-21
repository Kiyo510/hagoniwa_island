# hagoniwa_island
古のCGIゲームである箱庭諸島2を5系PHPからGolangへフルリプレイスするプロジェクトです。

## そもそも箱庭諸島って何？
- 箱庭諸島2は、徳岡宏樹が開発した、島を開発していくCGIゲームです。他にも数多く存在する箱庭諸島の元となった一般的なもので、派生バージョンである[究想の箱庭](https://hakopedia.uhyohyo.net/wiki/%E7%A9%B6%E6%83%B3%E3%81%AE%E7%AE%B1%E5%BA%AD)や[Hakoniwa R.A](https://hakopedia.uhyohyo.net/wiki/Hakoniwa_R.A.)などと比べると仕様がかなりシンプルで初心者向けの仕様となっています。
  - 参考: [箱庭諸島2 - Hakoniwapedia](https://hakopedia.uhyohyo.net/wiki/%E7%AE%B1%E5%BA%AD%E8%AB%B8%E5%B3%B62#:~:text=%E7%AE%B1%E5%BA%AD%E8%AB%B8%E5%B3%B62%E3%81%AF%E3%80%81%E5%BE%B3%E5%B2%A1,%E3%81%A7%E3%80%81%E5%88%9D%E5%BF%83%E8%80%85%E5%90%91%E3%81%91%E3%81%A7%E3%81%82%E3%82%8B%E3%80%82)

## 目的
- 箱庭諸島のスクリプトはPHP5系で書かれているが、PHP5系はサポートが終了しており保守もされていない。Golangへの移行を行いスクリプトを再配布することで、箱庭諸島をもっとたくさんの人に楽しんでもらう。
    - 実際、本番稼動している箱庭諸島2はわずかながら存在している。

## 方針
- [ざっくり方針](https://github.com/Kiyo510/hagoniwa_island/pull/18)
- 余裕があれば、最新のPHPバージョンでも動くようにする。(多分別プロジェクトにする)