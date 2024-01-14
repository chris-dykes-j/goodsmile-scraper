# Goodsmile Scraper

This project's purpose is to collect figure data from the Goodsmile website (https://www.goodsmile.info/en/).

### Current Goals:
- [ ] Transform json data for PostgreSQL database.
- [ ] Create database scheme for data.
- [ ] Load transformed data into dev database.
- [ ] Fix bug issues (continuous).

### Next Development Phase:
- [ ] Begin Web API rewrite. Include demo data and old data
- [ ] Build CLI Interaction
- [ ] Scrape by Year and/or Month. Goal to minimize traffic.
- [ ] Scrape specified section from command line, nendoroid or scale figure, etc.

### Completed Goals:
- [x] Collect nendoroid figure data. Store as JSON (structure is inconsistent).
- [x] Collect nendoroid images from site.
- [x] Collect data in English, Japanese, and Chinese.
