# Redis Memory Info Tests
## test1:
### key: [0-999] 
### val: 10000 bytes data
### set 1000 keys
### Result
(11432400 - 1120016) / 1000 = 10312.384  (3% extra)

### Memory Before Insert
used_memory:1120016
used_memory_human:1.07M
used_memory_rss:10407936
used_memory_rss_human:9.93M
used_memory_peak:8635392
used_memory_peak_human:8.24M
used_memory_peak_perc:12.97%
used_memory_overhead:1062560
used_memory_startup:1062560
used_memory_dataset:57456
used_memory_dataset_perc:100.00%
allocator_allocated:1063184
allocator_active:10370048
allocator_resident:10370048
total_system_memory:17179869184
total_system_memory_human:16.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:9.75
allocator_frag_bytes:9306864
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.00
rss_overhead_bytes:37888
mem_fragmentation_ratio:9.79
mem_fragmentation_bytes:9344752
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:0
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0


### Memory After Insert
used_memory:11432400
used_memory_human:10.90M
used_memory_rss:26304512
used_memory_rss_human:25.09M
used_memory_peak:89371440
used_memory_peak_human:85.23M
used_memory_peak_perc:12.79%
used_memory_overhead:1128192
used_memory_startup:1062560
used_memory_dataset:10304208
used_memory_dataset_perc:99.37%
allocator_allocated:1146832
allocator_active:26266624
allocator_resident:26266624
total_system_memory:17179869184
total_system_memory_human:16.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:22.90
allocator_frag_bytes:25119792
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.00
rss_overhead_bytes:37888
mem_fragmentation_ratio:22.94
mem_fragmentation_bytes:25157680
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

## test2:
### key: [0-999]
### val: 500000 bytes data
### set 1000 keys
### Result
(505487808 - 1120016) / 1000 = 504367.792 (0.8 % extra)

# Memory
used_memory:1120192
used_memory_human:1.07M
used_memory_rss:29851648
used_memory_rss_human:28.47M


# Memory
used_memory:505487808
used_memory_human:482.07M
used_memory_rss:527773696
used_memory_rss_human:503.32M


## test3:
### key: [0-999]
### val: 50000 bytes data
### set 1000 keys
### Result
(54506288 - 1120384) / 1000 = 53385.728 (6~7% extra)

### Memory Before
used_memory:1120384
used_memory_human:1.07M
used_memory_rss:26247168
used_memory_rss_human:25.03M

### Memory After
used_memory:54506288
used_memory_human:51.98M
used_memory_rss:64151552
used_memory_rss_human:61.18M

# Conclusion:
在忽略key的字节情况下，是存储的value平均字节的不超过7%的额外空间

