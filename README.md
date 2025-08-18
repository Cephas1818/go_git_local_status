# Go Git Local Status — Local Git Contribution Graph CLI Visualizer

[![Releases](https://img.shields.io/badge/Releases-Download-blue?style=for-the-badge&logo=github)](https://github.com/Cephas1818/go_git_local_status/releases)

A Go CLI that builds GitHub-style contribution graphs from local Git repos. View activity across many projects in the terminal. Track commit counts by day. Compare projects. Export data.

- Topics: cli, coding-activity, command-line, contribution-graph, developer-tools, development-tracking, git, git-analysis, git-history, git-stats, go, go-cli, golang, local-repositories, productivity, repository-analysis, terminal, terminal-app, visualization

Badges
- Build: ![Go](https://img.shields.io/badge/go-1.20+-00ADD8?logo=go)
- License: ![MIT](https://img.shields.io/badge/license-MIT-green)
- Releases: [https://github.com/Cephas1818/go_git_local_status/releases](https://github.com/Cephas1818/go_git_local_status/releases)

Screenshots

Terminal heatmap (example)
![Contribution graph example](https://raw.githubusercontent.com/github/explore/main/topics/graph/graph.png)

ASCII example output
```
Repo: ~/projects/api
Mon Tue Wed Thu Fri Sat Sun
 0   1   3   2   4   0   0
 1   0   0   2   5   1   0
 3   4   2   0   1   0   0
```

Why use this tool

- View local commit activity with the same mental model as GitHub contributions.
- Aggregate data from many repos in a single view.
- Run in terminals and scripts.
- Export CSV or JSON for reporting and dashboards.
- Use in dotfiles, CI reports, or daily dev metrics.

Install

1) Go install
```
go install github.com/Cephas1818/go_git_local_status@latest
```
After install, the binary will live in $GOPATH/bin or $HOME/go/bin. Add that to PATH if needed.

2) Download a release binary
Download the binary file for your platform from the Releases page and run it.
You must download and execute the asset file from:
https://github.com/Cephas1818/go_git_local_status/releases

Examples

Scan a single repo
```
go-git-local-status scan ~/projects/myrepo
```

Scan multiple repos
```
go-git-local-status scan ~/projects/*
```

Aggregate across a parent folder and show a 52-week heatmap
```
go-git-local-status heatmap --path ~/work --weeks 52
```

Export JSON
```
go-git-local-status export --path ~/work --format json --out commits.json
```

Show per-repo breakdown
```
go-git-local-status report --path ~/work --group-by repo
```

Commands and flags

- scan
  - --path string: Path to repo or folder.
  - --since date: Start date (YYYY-MM-DD).
  - --until date: End date (YYYY-MM-DD).
  - --depth int: Depth to scan subfolders.
- heatmap
  - --weeks int: Number of weeks to display.
  - --theme string: color|mono (terminal colors).
  - --min int: Minimum count for heat threshold.
- export
  - --format: json|csv
  - --out: Output file path.
- report
  - --group-by: repo|author|path
  - --top N: Show top N repos by commit count.

Output modes

- Terminal heatmap: Color squares or characters for each day.
- Plain ASCII: For terminals without color.
- CSV/JSON: Rows of date, repo, author, commits.
- Image export (SVG): Generates a static contribution graph SVG for embedding in docs.

Use cases

- Personal metrics: See your work pattern across many projects.
- Team snapshot: Collect commit activity across a monorepo or org folder.
- Backup reports: Export commit counts for audits.
- Embedding: Include a generated SVG in your personal site or README.
- Automation: Run in cron to create weekly reports.

How it works

- The tool runs git log in each repo folder.
- It parses commit date, author, and commit count.
- It aggregates commits per day, per repo.
- It maps counts to intensity levels for the heatmap.
- It renders a grid of weeks x days like GitHub contributions.
- It can export raw data for external analysis.

Performance notes

- The tool reads commit history. It scales with repo size.
- For large histories, use --since or --until to restrict range.
- Use --depth to limit scanned subfolders when scanning a parent folder.
- The tool spawns git subprocesses per repo. Run on multi-core systems for faster results.

Customization

- Theme: Choose color or mono. Terminal palette adapts to ANSI 256 colors.
- Thresholds: Set intensity buckets to change visual mapping.
- Grouping: Group by repo, author, or folder path.
- Output: Choose ASCII, color, SVG, JSON, or CSV.

Examples: real workflows

1) Weekly email snapshot
- Run a scan in CI each Sunday.
- Export JSON.
- Generate a small SVG.
- Attach SVG to an automated email.

2) Local dotfiles integration
- Add a cron job that saves the heatmap as SVG to ~/public.
- Show it on your personal site with a static image.

3) Team metrics
- Collect data from a shared workspace folder.
- Run report --group-by repo --top 10 to get the busiest projects.

Integration tips

- Use with tmux: The CLI fits in a pane. Resize to change layout.
- Shell prompt: Pipe the top N repos to a prompt helper.
- Git hooks: Add a post-commit hook to run a local scan and update a cache.

API and scripting

- The CLI returns exit codes for success and errors.
- Use --format json to parse output in scripts.
- Example: get today's total commits
```
go-git-local-status scan --path ~/work --since $(date +%F) --format json | jq '.total'
```

Internals (for contributors)

- Language: Go
- Core packages:
  - git: uses git CLI via os/exec for compatibility.
  - parser: parses git log output to structs.
  - aggregator: builds date -> count maps.
  - render: renders terminal, ascii, SVG, CSV, JSON.
- Design goals:
  - Small binary
  - Minimal runtime deps
  - Fast I/O and concurrent repo scanning

Contributing

- Fork the repo.
- Create a branch per feature or fix.
- Write tests for core parsing logic.
- Keep commits focused and small.
- Open a PR with a clear description and examples.

Release downloads

Download the binary file for your OS and CPU from the Releases page. After download, mark the file executable and run it. The Releases page is here:
https://github.com/Cephas1818/go_git_local_status/releases

License

- MIT

Security

- The tool runs git commands and reads repository history. It does not send data over the network by default.
- Review the code before running in sensitive environments.

Common issues and resolutions

- No output for a repo
  - Ensure the path contains a .git folder.
  - Run git status in the folder to confirm.
- Slow scan
  - Limit the date range with --since.
  - Use --depth to reduce the number of scanned subfolders.
- Color looks wrong
  - Switch to --theme mono or adjust terminal color settings.

FAQ

Q: Can I include uncommitted work?
A: No. The tool reads commit history. It cannot see uncommitted changes.

Q: Can I attribute commits by author email?
A: Yes. Use --group-by author and the tool will list per-author counts.

Q: Can I get a per-hour heatmap?
A: The core view focuses on daily counts. Export raw data and aggregate by hour externally.

Acknowledgements

- Inspired by GitHub's contribution graph.
- Uses the git CLI for history parsing.
- Uses Go for static binaries and speed.

Contact

- Report issues and feature requests on the repository Issues page.
- For release binaries and assets, download the file and run it from:
https://github.com/Cephas1818/go_git_local_status/releases

License: MIT