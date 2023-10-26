# UM32 implementation

This is ICFP 2006 um32 implementation in golang.

### Sandmark output


Machine `intel i7-1260P`
```bash
/usr/bin/time -v ./um32 sandmark.umz
....
SANDmark complete.
        Command being timed: "./um32 sandmark.umz"
        User time (seconds): 12.98
        System time (seconds): 0.17
        Percent of CPU this job got: 108%
        Elapsed (wall clock) time (h:mm:ss or m:ss): 0:12.13
        Average shared text size (kbytes): 0
        Average unshared data size (kbytes): 0
        Average stack size (kbytes): 0
        Average total size (kbytes): 0
        Maximum resident set size (kbytes): 11448
        Average resident set size (kbytes): 0
        Major (requiring I/O) page faults: 199
        Minor (reclaiming a frame) page faults: 47044
        Voluntary context switches: 10324
        Involuntary context switches: 404
        Swaps: 0
        File system inputs: 111
        File system outputs: 0
        Socket messages sent: 0
        Socket messages received: 0
        Signals delivered: 0
        Page size (bytes): 4096
        Exit status: 0
```

You don't need fancy structs for machine and instruction. Its simple machine that is encoded in sand platters by some cultist DUH.

Just GO and have some fun.

### Running UM

`go build` will create `um32` binary file

To unpack codex.umz you can do:
```
./um32 codex.umz | tee codex.um
```

Then open codex.um in your prefered editor and delete first few lines, you will see when to stop.

## Tasks/notes
### Intro:

Get familiar with UMIX, its commands, reading mails etc.

Figure out how to copy program into UMIX (umodem; similar to `cat <<EOF` )

### hack.bas

You are informed via mail that someone is misusing guest account. Try exploiting it more!

