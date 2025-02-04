package config

// DefaultConfigTemplate defines the configuration template for the EVM RPC configuration
const DefaultConfigTemplate = `
###############################################################################
###                             Cronos Configuration                        ###
###############################################################################

[memiavl]

# Enable defines if the memiavl should be enabled.
enable = {{ .MemIAVL.Enable }}

# ZeroCopy defines if the memiavl should return slices pointing to mmap-ed buffers directly (zero-copy),
# the zero-copied slices must not be retained beyond current block's execution.
zero-copy = {{ .MemIAVL.ZeroCopy }}

# AsyncCommitBuffer defines the size of asynchronous commit queue, this greatly improve block catching-up
# performance, -1 means synchronous commit.
async-commit-buffer = {{ .MemIAVL.AsyncCommitBuffer }}

# SnapshotKeepRecent defines what many old snapshots (excluding the latest one) to keep after new snapshots are taken.
snapshot-keep-recent = {{ .MemIAVL.SnapshotKeepRecent }}

# SnapshotInterval defines the block interval the memiavl snapshot is taken, default to 1000.
snapshot-interval = {{ .MemIAVL.SnapshotInterval }}

# CacheSize defines the size of the cache for each memiavl store, default to 1000.
cache-size = {{ .MemIAVL.CacheSize }}
`
