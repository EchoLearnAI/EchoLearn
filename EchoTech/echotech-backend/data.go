package main

// Auto‑generated mock question bank: 2 questions × 3 difficulties × 19 topics = 114 questions.
// IDs follow the pattern <topic-short>-<difficulty>-q<n> for easy (e), medium (m), hard (h).
// Feel free to tweak wording or add more questions later – the slice is ready for extension.

type Option struct {
	ID          string `json:"id"`
	Text        string `json:"text"`
	IsCorrect   bool   `json:"isCorrect"`
	Explanation string `json:"explanation"`
}

type Question struct {
	ID         string   `json:"id"`
	Topic      string   `json:"topic"`
	Difficulty string   `json:"difficulty"`
	Text       string   `json:"text"`
	Options    []Option `json:"options"`
}

var questions = []Question{
	// ──────────────────────── Network ────────────────────────
	{
		ID:         "net-e-q1",
		Topic:      "Network",
		Difficulty: "easy",
		Text:       "What does TCP stand for?",
		Options: []Option{
			{ID: "opt1", Text: "Transmission Control Protocol", IsCorrect: true, Explanation: "Correct. TCP provides reliable, ordered delivery of data packets."},
			{ID: "opt2", Text: "Total Connection Process", IsCorrect: false, Explanation: "Not an actual networking term."},
			{ID: "opt3", Text: "Transfer Control Protocol", IsCorrect: false, Explanation: "Close, but the first word should be Transmission."},
			{ID: "opt4", Text: "Technical Communication Protocol", IsCorrect: false, Explanation: "Invented distractor."},
		},
	},
	{
		ID:         "net-e-q2",
		Topic:      "Network",
		Difficulty: "easy",
		Text:       "Which OSI layer is responsible for routing packets between networks?",
		Options: []Option{
			{ID: "opt1", Text: "Network layer (Layer 3)", IsCorrect: true, Explanation: "Correct. Layer 3 handles logical addressing and routing."},
			{ID: "opt2", Text: "Data Link layer (Layer 2)", IsCorrect: false, Explanation: "Layer 2 deals with framing and physical addressing within a LAN."},
			{ID: "opt3", Text: "Transport layer (Layer 4)", IsCorrect: false, Explanation: "Layer 4 handles end‑to‑end transport services like TCP/UDP."},
			{ID: "opt4", Text: "Session layer (Layer 5)", IsCorrect: false, Explanation: "Layer 5 establishes, manages, and terminates sessions."},
		},
	},
	{
		ID:         "net-m-q1",
		Topic:      "Network",
		Difficulty: "medium",
		Text:       "Which port number is used by HTTPS by default?",
		Options: []Option{
			{ID: "opt1", Text: "443", IsCorrect: true, Explanation: "Correct. HTTPS typically operates over TCP port 443."},
			{ID: "opt2", Text: "80", IsCorrect: false, Explanation: "Port 80 is used by unencrypted HTTP."},
			{ID: "opt3", Text: "22", IsCorrect: false, Explanation: "Port 22 is used by SSH."},
			{ID: "opt4", Text: "8080", IsCorrect: false, Explanation: "8080 is a common alternative HTTP port, not HTTPS."},
		},
	},
	{
		ID:         "net-m-q2",
		Topic:      "Network",
		Difficulty: "medium",
		Text:       "In IPv6 addressing, how many bits are in the address?",
		Options: []Option{
			{ID: "opt1", Text: "128", IsCorrect: true, Explanation: "Correct. IPv6 addresses are 128 bits long."},
			{ID: "opt2", Text: "32", IsCorrect: false, Explanation: "32‑bit addresses are used in IPv4."},
			{ID: "opt3", Text: "64", IsCorrect: false, Explanation: "64‑bit is half of an IPv6 address, typically the interface identifier length."},
			{ID: "opt4", Text: "256", IsCorrect: false, Explanation: "IPv6 does not use 256‑bit addresses."},
		},
	},
	{
		ID:         "net-h-q1",
		Topic:      "Network",
		Difficulty: "hard",
		Text:       "Which TCP flag combination is found in the second step of the three‑way handshake?",
		Options: []Option{
			{ID: "opt1", Text: "SYN + ACK", IsCorrect: true, Explanation: "Correct. The server replies to SYN with a SYN+ACK."},
			{ID: "opt2", Text: "ACK only", IsCorrect: false, Explanation: "An ACK alone is not used until the third step."},
			{ID: "opt3", Text: "FIN + ACK", IsCorrect: false, Explanation: "FIN is used to terminate a connection, not establish it."},
			{ID: "opt4", Text: "PSH + ACK", IsCorrect: false, Explanation: "PSH is used to push buffered data through, not for handshakes."},
		},
	},
	{
		ID:         "net-h-q2",
		Topic:      "Network",
		Difficulty: "hard",
		Text:       "Which routing protocol uses link‑state advertisements and Dijkstra's algorithm to build its routing table?",
		Options: []Option{
			{ID: "opt1", Text: "OSPF", IsCorrect: true, Explanation: "Correct. OSPF is a link‑state protocol that uses Dijkstra to compute shortest paths."},
			{ID: "opt2", Text: "RIP", IsCorrect: false, Explanation: "RIP is a distance‑vector protocol using hop count and Bellman‑Ford."},
			{ID: "opt3", Text: "BGP", IsCorrect: false, Explanation: "BGP is a path‑vector protocol used between autonomous systems."},
			{ID: "opt4", Text: "EIGRP", IsCorrect: false, Explanation: "EIGRP uses a diffusing update algorithm (DUAL), not pure Dijkstra."},
		},
	},

	// ──────────────────────── Cybersecurity ────────────────────────
	{
		ID:         "sec-e-q1",
		Topic:      "Cybersecurity",
		Difficulty: "easy",
		Text:       "What does the acronym CIA stand for in security?",
		Options: []Option{
			{ID: "opt1", Text: "Confidentiality, Integrity, Availability", IsCorrect: true, Explanation: "Correct. CIA triad represents the pillars of information security."},
			{ID: "opt2", Text: "Central Intelligence Agency", IsCorrect: false, Explanation: "In security context, CIA refers to a triad, not the U.S. agency."},
			{ID: "opt3", Text: "Control, Inspection, Audit", IsCorrect: false, Explanation: "These do not form the standard triad."},
			{ID: "opt4", Text: "Confidentiality, Inspection, Authorization", IsCorrect: false, Explanation: "Integrity and Availability are missing."},
		},
	},
	{
		ID:         "sec-e-q2",
		Topic:      "Cybersecurity",
		Difficulty: "easy",
		Text:       "Which mechanism proves a user's identity?",
		Options: []Option{
			{ID: "opt1", Text: "Authentication", IsCorrect: true, Explanation: "Correct. Authentication verifies identity."},
			{ID: "opt2", Text: "Authorization", IsCorrect: false, Explanation: "Authorization controls what an authenticated user can do."},
			{ID: "opt3", Text: "Accounting", IsCorrect: false, Explanation: "Accounting tracks user activity."},
			{ID: "opt4", Text: "Auditing", IsCorrect: false, Explanation: "Auditing reviews logs and controls, not identity proof."},
		},
	},
	{
		ID:         "sec-m-q1",
		Topic:      "Cybersecurity",
		Difficulty: "medium",
		Text:       "Which attack technique exploits input fields by injecting malicious SQL statements?",
		Options: []Option{
			{ID: "opt1", Text: "SQL Injection", IsCorrect: true, Explanation: "Correct. Malicious SQL is inserted into queries."},
			{ID: "opt2", Text: "Cross‑Site Scripting", IsCorrect: false, Explanation: "XSS injects scripts into web pages, not SQL."},
			{ID: "opt3", Text: "CSRF", IsCorrect: false, Explanation: "CSRF forces users to execute unwanted actions."},
			{ID: "opt4", Text: "Buffer Overflow", IsCorrect: false, Explanation: "Overwrites memory, unrelated to SQL syntax."},
		},
	},
	{
		ID:         "sec-m-q2",
		Topic:      "Cybersecurity",
		Difficulty: "medium",
		Text:       "Which cryptographic algorithm is asymmetric?",
		Options: []Option{
			{ID: "opt1", Text: "RSA", IsCorrect: true, Explanation: "Correct. RSA uses a public/private key pair."},
			{ID: "opt2", Text: "AES", IsCorrect: false, Explanation: "AES is symmetric."},
			{ID: "opt3", Text: "ChaCha20", IsCorrect: false, Explanation: "ChaCha20 is symmetric."},
			{ID: "opt4", Text: "DES", IsCorrect: false, Explanation: "DES is symmetric."},
		},
	},
	{
		ID:         "sec-h-q1",
		Topic:      "Cybersecurity",
		Difficulty: "hard",
		Text:       "Which tool implements the Kerberoasting attack to retrieve service tickets?",
		Options: []Option{
			{ID: "opt1", Text: "Impacket's GetUserSPNs.py", IsCorrect: true, Explanation: "Correct. The script extracts RC4‑encrypted service tickets for offline cracking."},
			{ID: "opt2", Text: "John the Ripper", IsCorrect: false, Explanation: "John can crack hashes but doesn't extract service tickets."},
			{ID: "opt3", Text: "Nmap", IsCorrect: false, Explanation: "Nmap is a scanner, not used for Kerberoasting."},
			{ID: "opt4", Text: "Aircrack‑ng", IsCorrect: false, Explanation: "Aircrack‑ng targets Wi‑Fi, not Kerberos."},
		},
	},
	{
		ID:         "sec-h-q2",
		Topic:      "Cybersecurity",
		Difficulty: "hard",
		Text:       "Which security model introduces 'no read up, no write down' rules?",
		Options: []Option{
			{ID: "opt1", Text: "Bell‑LaPadula", IsCorrect: true, Explanation: "Correct. Bell‑LaPadula enforces confidentiality with these rules."},
			{ID: "opt2", Text: "Biba", IsCorrect: false, Explanation: "Biba focuses on integrity: no write up, no read down."},
			{ID: "opt3", Text: "Clark‑Wilson", IsCorrect: false, Explanation: "Clark‑Wilson enforces well‑formed transactions for integrity."},
			{ID: "opt4", Text: "Chinese Wall", IsCorrect: false, Explanation: "Chinese Wall prevents conflicts of interest."},
		},
	},

	// ──────────────────────── VMs ────────────────────────
	{
		ID:         "vm-e-q1",
		Topic:      "VMs",
		Difficulty: "easy",
		Text:       "Which component emulates hardware for guest operating systems?",
		Options: []Option{
			{ID: "opt1", Text: "Hypervisor", IsCorrect: true, Explanation: "Correct. The hypervisor abstracts and manages hardware resources."},
			{ID: "opt2", Text: "Container runtime", IsCorrect: false, Explanation: "Containers share the host OS kernel rather than emulating hardware."},
			{ID: "opt3", Text: "Orchestrator", IsCorrect: false, Explanation: "Orchestrators schedule containers/VMs; they don't emulate hardware."},
			{ID: "opt4", Text: "Firmware", IsCorrect: false, Explanation: "Firmware is code embedded in hardware components."},
		},
	},
	{
		ID:         "vm-e-q2",
		Topic:      "VMs",
		Difficulty: "easy",
		Text:       "VMware Workstation is an example of which type of hypervisor?",
		Options: []Option{
			{ID: "opt1", Text: "Type‑2", IsCorrect: true, Explanation: "Correct. It runs on top of an existing OS."},
			{ID: "opt2", Text: "Type‑1", IsCorrect: false, Explanation: "Type‑1 hypervisors run directly on hardware (e.g., ESXi)."},
			{ID: "opt3", Text: "Bare‑metal", IsCorrect: false, Explanation: "Bare‑metal is another term for Type‑1, not applicable here."},
			{ID: "opt4", Text: "Para‑virtual", IsCorrect: false, Explanation: "Para‑virtualization is a technique, not a hypervisor type category here."},
		},
	},
	{
		ID:         "vm-m-q1",
		Topic:      "VMs",
		Difficulty: "medium",
		Text:       "Which CPU feature enables a VM to run instructions directly on hardware without trapping to the hypervisor?",
		Options: []Option{
			{ID: "opt1", Text: "Hardware virtualization extensions (Intel VT‑x / AMD‑V)", IsCorrect: true, Explanation: "Correct. VT‑x and AMD‑V allow safe direct execution."},
			{ID: "opt2", Text: "Hyper‑threading", IsCorrect: false, Explanation: "Hyper‑threading duplicates CPU resources for threads."},
			{ID: "opt3", Text: "SpeedStep", IsCorrect: false, Explanation: "SpeedStep manages CPU power, unrelated to virtualization."},
			{ID: "opt4", Text: "NUMA", IsCorrect: false, Explanation: "Non‑Uniform Memory Access relates to memory topology."},
		},
	},
	{
		ID:         "vm-m-q2",
		Topic:      "VMs",
		Difficulty: "medium",
		Text:       "What is live migration in virtualization platforms?",
		Options: []Option{
			{ID: "opt1", Text: "Moving a running VM to another host with minimal downtime", IsCorrect: true, Explanation: "Correct. Live migration keeps services running during relocation."},
			{ID: "opt2", Text: "Cloning a VM while it is powered off", IsCorrect: false, Explanation: "That's cold cloning, not live migration."},
			{ID: "opt3", Text: "Creating a snapshot of VM memory", IsCorrect: false, Explanation: "Snapshots capture state but don't relocate VMs."},
			{ID: "opt4", Text: "Converting a VM to a template", IsCorrect: false, Explanation: "Templates are for deployment, unrelated to migration."},
		},
	},
	{
		ID:         "vm-h-q1",
		Topic:      "VMs",
		Difficulty: "hard",
		Text:       "Which QEMU feature allows device emulation to run directly on the host for near‑native performance?",
		Options: []Option{
			{ID: "opt1", Text: "VFIO passthrough", IsCorrect: true, Explanation: "Correct. VFIO passes devices directly to the guest."},
			{ID: "opt2", Text: "KSM (Kernel Same‑page Merging)", IsCorrect: false, Explanation: "KSM deduplicates memory pages, not device emulation."},
			{ID: "opt3", Text: "VirtIO serial", IsCorrect: false, Explanation: "VirtIO serial is a paravirtualized device, not a passthrough."},
			{ID: "opt4", Text: "Spice", IsCorrect: false, Explanation: "Spice provides remote desktop, not device passthrough."},
		},
	},
	{
		ID:         "vm-h-q2",
		Topic:      "VMs",
		Difficulty: "hard",
		Text:       "In KVM, which structure tracks shadow page tables for guest memory virtualization?",
		Options: []Option{
			{ID: "opt1", Text: "MMU (Memory Management Unit) role", IsCorrect: true, Explanation: "Correct. KVM maintains shadow page tables in concert with host MMU."},
			{ID: "opt2", Text: "EPT (Extended Page Tables) only", IsCorrect: false, Explanation: "EPT is Intel's hardware feature; the structure managing them is shadowed."},
			{ID: "opt3", Text: "TLB", IsCorrect: false, Explanation: "TLB caches translations but isn't the shadow table structure."},
			{ID: "opt4", Text: "RAM disk", IsCorrect: false, Explanation: "RAM disk is unrelated to paging structures."},
		},
	},

	// ──────────────────────── Container ────────────────────────
	{
		ID:         "ctr-e-q1",
		Topic:      "Container",
		Difficulty: "easy",
		Text:       "Which command starts a new container in Docker?",
		Options: []Option{
			{ID: "opt1", Text: "docker run", IsCorrect: true, Explanation: "Correct. 'docker run' creates and starts a container."},
			{ID: "opt2", Text: "docker build", IsCorrect: false, Explanation: "'docker build' creates an image, not a container."},
			{ID: "opt3", Text: "docker pull", IsCorrect: false, Explanation: "'docker pull' downloads an image."},
			{ID: "opt4", Text: "docker ps", IsCorrect: false, Explanation: "'docker ps' lists running containers."},
		},
	},
	{
		ID:         "ctr-e-q2",
		Topic:      "Container",
		Difficulty: "easy",
		Text:       "Which Docker command lists running containers?",
		Options: []Option{
			{ID: "opt1", Text: "docker ps", IsCorrect: true, Explanation: "Correct. 'docker ps' shows running containers."},
			{ID: "opt2", Text: "docker images", IsCorrect: false, Explanation: "'docker images' lists local images."},
			{ID: "opt3", Text: "docker stop", IsCorrect: false, Explanation: "'docker stop' halts a running container."},
			{ID: "opt4", Text: "docker build", IsCorrect: false, Explanation: "'docker build' creates an image from a Dockerfile."},
		},
	},
	{
		ID:         "ctr-m-q1",
		Topic:      "Container",
		Difficulty: "medium",
		Text:       "Which Linux kernel feature enforces resource limits for containers?",
		Options: []Option{
			{ID: "opt1", Text: "cgroups", IsCorrect: true, Explanation: "Correct. Control groups limit CPU, memory, I/O, etc."},
			{ID: "opt2", Text: "iptables", IsCorrect: false, Explanation: "iptables manages packet filtering, not resource limits."},
			{ID: "opt3", Text: "systemd-units", IsCorrect: false, Explanation: "systemd runs services, but resource limits are in cgroups."},
			{ID: "opt4", Text: "SELinux", IsCorrect: false, Explanation: "SELinux enforces mandatory access control, not resource quotas."},
		},
	},
	{
		ID:         "ctr-m-q2",
		Topic:      "Container",
		Difficulty: "medium",
		Text:       "Which OCI-compliant runtime does Docker use by default under the hood?",
		Options: []Option{
			{ID: "opt1", Text: "runc", IsCorrect: true, Explanation: "Correct. Docker’s default low-level runtime is runc."},
			{ID: "opt2", Text: "containerd", IsCorrect: false, Explanation: "containerd is a higher-level daemon that calls runc."},
			{ID: "opt3", Text: "cri-o", IsCorrect: false, Explanation: "CRI-O is a lightweight runtime for Kubernetes, not Docker."},
			{ID: "opt4", Text: "Kata Containers", IsCorrect: false, Explanation: "Kata is an alternative runtime focusing on hardware isolation."},
		},
	},
	{
		ID:         "ctr-h-q1",
		Topic:      "Container",
		Difficulty: "hard",
		Text:       "What does the 'USER nobody' instruction in a Dockerfile mitigate?",
		Options: []Option{
			{ID: "opt1", Text: "Running the container as root", IsCorrect: true, Explanation: "Correct. It drops privileges inside the image."},
			{ID: "opt2", Text: "Image size bloat", IsCorrect: false, Explanation: "USER does not affect layers or size."},
			{ID: "opt3", Text: "Network namespace conflicts", IsCorrect: false, Explanation: "USER is unrelated to namespaces selection."},
			{ID: "opt4", Text: "CPU throttling", IsCorrect: false, Explanation: "CPU quotas are set with cgroups, not USER."},
		},
	},
	{
		ID:         "ctr-h-q2",
		Topic:      "Container",
		Difficulty: "hard",
		Text:       "Which flag to 'docker run' enables an AppArmor profile for extra confinement?",
		Options: []Option{
			{ID: "opt1", Text: "--security-opt", IsCorrect: true, Explanation: "Correct. AppArmor or seccomp can be specified via --security-opt."},
			{ID: "opt2", Text: "--privileged", IsCorrect: false, Explanation: "This broadens privileges, not confines them."},
			{ID: "opt3", Text: "--network host", IsCorrect: false, Explanation: "This puts the container in the host network namespace."},
			{ID: "opt4", Text: "--detach", IsCorrect: false, Explanation: "Runs the container in the background but has no security effect."},
		},
	},

	// ──────────────────────── Architecture ────────────────────────
	{
		ID:         "arch-e-q1",
		Topic:      "Architecture",
		Difficulty: "easy",
		Text:       "Which architectural style advocates building small, independently deployable services?",
		Options: []Option{
			{ID: "opt1", Text: "Microservices", IsCorrect: true, Explanation: "Correct. Microservices decompose a system into loosely coupled services."},
			{ID: "opt2", Text: "Monolith", IsCorrect: false, Explanation: "Monoliths are single-deployment units."},
			{ID: "opt3", Text: "Batch processing", IsCorrect: false, Explanation: "Batch is about processing mode, not service size."},
			{ID: "opt4", Text: "FTP", IsCorrect: false, Explanation: "FTP is a protocol, not an architecture style."},
		},
	},
	{
		ID:         "arch-e-q2",
		Topic:      "Architecture",
		Difficulty: "easy",
		Text:       "In a client-server architecture, the server provides ___?",
		Options: []Option{
			{ID: "opt1", Text: "Resources or services requested by the client", IsCorrect: true, Explanation: "Correct. Clients consume; servers serve."},
			{ID: "opt2", Text: "DNS name resolution", IsCorrect: false, Explanation: "DNS is a separate service."},
			{ID: "opt3", Text: "Only user interfaces", IsCorrect: false, Explanation: "UIs run on client side or web front ends."},
			{ID: "opt4", Text: "Compilation of code", IsCorrect: false, Explanation: "Build servers compile, but it's not inherent to client-server."},
		},
	},
	{
		ID:         "arch-m-q1",
		Topic:      "Architecture",
		Difficulty: "medium",
		Text:       "Which two guarantees can a distributed system choose simultaneously according to CAP theorem?",
		Options: []Option{
			{ID: "opt1", Text: "Consistency and partition tolerance, or availability and partition tolerance", IsCorrect: true, Explanation: "Correct. Under a partition you must trade A vs C."},
			{ID: "opt2", Text: "Consistency and availability always", IsCorrect: false, Explanation: "Impossible when a network partition occurs."},
			{ID: "opt3", Text: "Partition tolerance alone", IsCorrect: false, Explanation: "One other guarantee must accompany P."},
			{ID: "opt4", Text: "Atomicity and durability", IsCorrect: false, Explanation: "AD are ACID properties, not CAP."},
		},
	},
	{
		ID:         "arch-m-q2",
		Topic:      "Architecture",
		Difficulty: "medium",
		Text:       "What does CQRS stand for?",
		Options: []Option{
			{ID: "opt1", Text: "Command Query Responsibility Segregation", IsCorrect: true, Explanation: "Correct. CQRS separates read and write models."},
			{ID: "opt2", Text: "Convergent Query Replication Service", IsCorrect: false, Explanation: "Made-up phrase."},
			{ID: "opt3", Text: "Centralized Queue Routing System", IsCorrect: false, Explanation: "Made-up phrase."},
			{ID: "opt4", Text: "Continuous Quality Review Session", IsCorrect: false, Explanation: "Not an architectural acronym."},
		},
	},
	{
		ID:         "arch-h-q1",
		Topic:      "Architecture",
		Difficulty: "hard",
		Text:       "Which pattern ensures that only one instance of a service runs active while another stays hot-standby?",
		Options: []Option{
			{ID: "opt1", Text: "Active-passive failover", IsCorrect: true, Explanation: "Correct. One active, one standby."},
			{ID: "opt2", Text: "Active-active clustering", IsCorrect: false, Explanation: "Both nodes serve traffic in active-active."},
			{ID: "opt3", Text: "Round-robin load balancing", IsCorrect: false, Explanation: "This spreads load across many actives."},
			{ID: "opt4", Text: "Sharding", IsCorrect: false, Explanation: "Sharding splits data by key, not hot-standby."},
		},
	},
	{
		ID:         "arch-h-q2",
		Topic:      "Architecture",
		Difficulty: "hard",
		Text:       "In event-driven architecture, what is the primary benefit of using an event sourcing model?",
		Options: []Option{
			{ID: "opt1", Text: "It stores every state change as an immutable event", IsCorrect: true, Explanation: "Correct. The full event log reconstructs state anytime."},
			{ID: "opt2", Text: "It eliminates the need for a database", IsCorrect: false, Explanation: "You still need durable storage—the event log itself."},
			{ID: "opt3", Text: "It reduces network latency to zero", IsCorrect: false, Explanation: "Latency still exists."},
			{ID: "opt4", Text: "It disallows scaling consumers", IsCorrect: false, Explanation: "Event streams are actually very scalable."},
		},
	},

	// ──────────────────────── Databases ────────────────────────
	{
		ID:         "db-e-q1",
		Topic:      "Databases",
		Difficulty: "easy",
		Text:       "Which SQL clause filters rows based on a condition?",
		Options: []Option{
			{ID: "opt1", Text: "WHERE", IsCorrect: true, Explanation: "Correct. WHERE restricts rows returned."},
			{ID: "opt2", Text: "SELECT", IsCorrect: false, Explanation: "SELECT lists columns to retrieve."},
			{ID: "opt3", Text: "GROUP BY", IsCorrect: false, Explanation: "GROUP BY aggregates rows."},
			{ID: "opt4", Text: "ORDER BY", IsCorrect: false, Explanation: "ORDER BY sorts the final result set."},
		},
	},
	{
		ID:         "db-e-q2",
		Topic:      "Databases",
		Difficulty: "easy",
		Text:       "MySQL, PostgreSQL, and MariaDB are examples of ___ databases.",
		Options: []Option{
			{ID: "opt1", Text: "Relational", IsCorrect: true, Explanation: "Correct. They use tables with relational integrity."},
			{ID: "opt2", Text: "Graph", IsCorrect: false, Explanation: "Graph DBs focus on nodes/edges (e.g., Neo4j)."},
			{ID: "opt3", Text: "Key-value", IsCorrect: false, Explanation: "Examples include Redis, etcd."},
			{ID: "opt4", Text: "Time-series", IsCorrect: false, Explanation: "Examples include InfluxDB, Prometheus TSDB."},
		},
	},
	{
		ID:         "db-m-q1",
		Topic:      "Databases",
		Difficulty: "medium",
		Text:       "Which isolation level in SQL prevents dirty reads but allows non-repeatable reads?",
		Options: []Option{
			{ID: "opt1", Text: "Read committed", IsCorrect: true, Explanation: "Correct. Read committed disallows dirty reads."},
			{ID: "opt2", Text: "Repeatable read", IsCorrect: false, Explanation: "Repeatable read disallows non-repeatable reads too."},
			{ID: "opt3", Text: "Serializable", IsCorrect: false, Explanation: "Serializable is the strictest isolation."},
			{ID: "opt4", Text: "Read uncommitted", IsCorrect: false, Explanation: "Read uncommitted allows dirty reads."},
		},
	},
	{
		ID:         "db-m-q2",
		Topic:      "Databases",
		Difficulty: "medium",
		Text:       "In PostgreSQL, which command creates a new database user?",
		Options: []Option{
			{ID: "opt1", Text: "CREATE ROLE", IsCorrect: true, Explanation: "Correct. CREATE ROLE with LOGIN creates a user."},
			{ID: "opt2", Text: "ADD USER", IsCorrect: false, Explanation: "Not valid SQL syntax in Postgres."},
			{ID: "opt3", Text: "INSERT USER", IsCorrect: false, Explanation: "INSERT is for data rows."},
			{ID: "opt4", Text: "ALTER DATABASE", IsCorrect: false, Explanation: "Alters DB properties, not users."},
		},
	},
	{
		ID:         "db-h-q1",
		Topic:      "Databases",
		Difficulty: "hard",
		Text:       "Which consistency model does Cassandra provide by default?",
		Options: []Option{
			{ID: "opt1", Text: "Eventual consistency", IsCorrect: true, Explanation: "Correct. Tunable consistency per query, but eventual by default."},
			{ID: "opt2", Text: "Strong consistency", IsCorrect: false, Explanation: "Cassandra can approach strongconsistency with quorum, but not default."},
			{ID: "opt3", Text: "Linearizability", IsCorrect: false, Explanation: "Etcd/Raft systems are linearizable, Cassandra is not."},
			{ID: "opt4", Text: "Snapshot isolation", IsCorrect: false, Explanation: "Snapshot isolation is an RDBMS term."},
		},
	},
	{
		ID:         "db-h-q2",
		Topic:      "Databases",
		Difficulty: "hard",
		Text:       "What data structure underpins an InnoDB secondary index?",
		Options: []Option{
			{ID: "opt1", Text: "B+ tree", IsCorrect: true, Explanation: "Correct. Both clustered and secondary indexes use B+ trees."},
			{ID: "opt2", Text: "Hash map", IsCorrect: false, Explanation: "MySQL MEMORY engine offers HASH, not InnoDB on-disk indexes."},
			{ID: "opt3", Text: "Skip list", IsCorrect: false, Explanation: "Skip lists are used in Redis sorted sets."},
			{ID: "opt4", Text: "Trie", IsCorrect: false, Explanation: "Tries store prefixes, not typical for InnoDB."},
		},
	},

	// ──────────────────────── Git ────────────────────────
	{
		ID:         "git-e-q1",
		Topic:      "Git",
		Difficulty: "easy",
		Text:       "Which command stages changes for the next commit?",
		Options: []Option{
			{ID: "opt1", Text: "git add", IsCorrect: true, Explanation: "Correct. Adds files to the index."},
			{ID: "opt2", Text: "git commit -m", IsCorrect: false, Explanation: "Commits what's already staged."},
			{ID: "opt3", Text: "git push", IsCorrect: false, Explanation: "Pushes commits to a remote repository."},
			{ID: "opt4", Text: "git status", IsCorrect: false, Explanation: "Shows current working tree status."},
		},
	},
	{
		ID:         "git-e-q2",
		Topic:      "Git",
		Difficulty: "easy",
		Text:       "Which file tells Git which files or patterns to ignore?",
		Options: []Option{
			{ID: "opt1", Text: ".gitignore", IsCorrect: true, Explanation: "Correct. Lines in .gitignore exclude files from commits."},
			{ID: "opt2", Text: ".gitconfig", IsCorrect: false, Explanation: "gitconfig stores configuration, not ignore rules."},
			{ID: "opt3", Text: "README.md", IsCorrect: false, Explanation: "README is for documentation."},
			{ID: "opt4", Text: ".gitattributes", IsCorrect: false, Explanation: "gitattributes handles text conversion and filters."},
		},
	},
	{
		ID:         "git-m-q1",
		Topic:      "Git",
		Difficulty: "medium",
		Text:       "Which option to 'git log' shows each commit on a single line and a graph?",
		Options: []Option{
			{ID: "opt1", Text: "--oneline --graph", IsCorrect: true, Explanation: "Correct. Combines concise output with ASCII graph."},
			{ID: "opt2", Text: "--stat", IsCorrect: false, Explanation: "--stat shows diffstats, not graph."},
			{ID: "opt3", Text: "--patch", IsCorrect: false, Explanation: "--patch includes diffs."},
			{ID: "opt4", Text: "--name-only", IsCorrect: false, Explanation: "Shows file names changed, no graph."},
		},
	},
	{
		ID:         "git-m-q2",
		Topic:      "Git",
		Difficulty: "medium",
		Text:       "After cloning, which command downloads new commits from origin without merging?",
		Options: []Option{
			{ID: "opt1", Text: "git fetch", IsCorrect: true, Explanation: "Correct. fetch retrieves but doesn't integrate changes."},
			{ID: "opt2", Text: "git pull", IsCorrect: false, Explanation: "pull fetches and merges/rebases."},
			{ID: "opt3", Text: "git remote add", IsCorrect: false, Explanation: "Adds remote but doesn't download."},
			{ID: "opt4", Text: "git checkout", IsCorrect: false, Explanation: "Switches branches or restores files."},
		},
	},
	{
		ID:         "git-h-q1",
		Topic:      "Git",
		Difficulty: "hard",
		Text:       "Which command rewrites history by creating new commits without changing working files?",
		Options: []Option{
			{ID: "opt1", Text: "git rebase --interactive", IsCorrect: true, Explanation: "Correct. Rebase can squash, reword, reorder commits."},
			{ID: "opt2", Text: "git merge --no-ff", IsCorrect: false, Explanation: "Merge creates a merge commit but doesn't rewrite existing commits."},
			{ID: "opt3", Text: "git cherry-pick", IsCorrect: false, Explanation: "Cherry-pick copies a commit but doesn't rewrite history of branch."},
			{ID: "opt4", Text: "git tag", IsCorrect: false, Explanation: "Tags annotate commits but don't rewrite."},
		},
	},
	{
		ID:         "git-h-q2",
		Topic:      "Git",
		Difficulty: "hard",
		Text:       "What does 'git rev-parse HEAD^@' output?",
		Options: []Option{
			{ID: "opt1", Text: "The full SHA-1 hashes of each parent of HEAD", IsCorrect: true, Explanation: "Correct. '^@' expands to all parents."},
			{ID: "opt2", Text: "The tree object of the HEAD commit", IsCorrect: false, Explanation: "That would be 'git rev-parse HEAD^{tree}'."},
			{ID: "opt3", Text: "The previous commit’s hash only", IsCorrect: false, Explanation: "HEAD^ yields the first parent, not ^@."},
			{ID: "opt4", Text: "A list of branches pointing to HEAD", IsCorrect: false, Explanation: "Branches are refs, not returned by rev-parse."},
		},
	},

	// ──────────────────────── Cache & CDN ────────────────────────
	{
		ID:         "cache-e-q1",
		Topic:      "Cache and CDN",
		Difficulty: "easy",
		Text:       "Which HTTP header controls how long a response may be cached?",
		Options: []Option{
			{ID: "opt1", Text: "Cache-Control", IsCorrect: true, Explanation: "Correct. Directives such as max-age define lifetime."},
			{ID: "opt2", Text: "Content-Type", IsCorrect: false, Explanation: "Specifies MIME type."},
			{ID: "opt3", Text: "ETag", IsCorrect: false, Explanation: "ETag identifies versions but doesn't set lifetime."},
			{ID: "opt4", Text: "Set-Cookie", IsCorrect: false, Explanation: "Cookies transmit session data, not caching policy."},
		},
	},
	{
		ID:         "cache-e-q2",
		Topic:      "Cache and CDN",
		Difficulty: "easy",
		Text:       "A CDN primarily improves ___ for globally distributed users.",
		Options: []Option{
			{ID: "opt1", Text: "Latency", IsCorrect: true, Explanation: "Correct. Edge nodes serve content nearer to users."},
			{ID: "opt2", Text: "Server CPU utilization only", IsCorrect: false, Explanation: "CPU relief is secondary; goal is lower latency."},
			{ID: "opt3", Text: "Encryption strength", IsCorrect: false, Explanation: "Encryption unaffected by CDN."},
			{ID: "opt4", Text: "Programming language choice", IsCorrect: false, Explanation: "CDN is independent of code language."},
		},
	},
	{
		ID:         "cache-m-q1",
		Topic:      "Cache and CDN",
		Difficulty: "medium",
		Text:       "Redis uses which eviction policy by default when memory is full and policy is not set?",
		Options: []Option{
			{ID: "opt1", Text: "No eviction (returns an error)", IsCorrect: true, Explanation: "Correct. Policy 'noeviction' returns error on writes."},
			{ID: "opt2", Text: "LRU eviction", IsCorrect: false, Explanation: "LRU requires 'allkeys-lru' or similar policy set."},
			{ID: "opt3", Text: "LFU eviction", IsCorrect: false, Explanation: "LFU is 'allkeys-lfu'."},
			{ID: "opt4", Text: "Random eviction", IsCorrect: false, Explanation: "Random is 'allkeys-random'."},
		},
	},
	{
		ID:         "cache-m-q2",
		Topic:      "Cache and CDN",
		Difficulty: "medium",
		Text:       "Which DNS record type is commonly used to map a CDN hostname to an origin?",
		Options: []Option{
			{ID: "opt1", Text: "CNAME", IsCorrect: true, Explanation: "Correct. Alias record points to CDN domain."},
			{ID: "opt2", Text: "A", IsCorrect: false, Explanation: "A maps directly to an IP; CDNs abstract this."},
			{ID: "opt3", Text: "TXT", IsCorrect: false, Explanation: "TXT holds arbitrary text, often SPF or verification."},
			{ID: "opt4", Text: "MX", IsCorrect: false, Explanation: "MX is for mail exchangers."},
		},
	},
	{
		ID:         "cache-h-q1",
		Topic:      "Cache and CDN",
		Difficulty: "hard",
		Text:       "In Varnish Configuration Language (VCL), which subroutine decides how to fetch content from the backend?",
		Options: []Option{
			{ID: "opt1", Text: "vcl_backend_fetch", IsCorrect: true, Explanation: "Correct. It controls backend request parameters."},
			{ID: "opt2", Text: "vcl_recv", IsCorrect: false, Explanation: "Called when a request is received."},
			{ID: "opt3", Text: "vcl_deliver", IsCorrect: false, Explanation: "Runs before response is delivered."},
			{ID: "opt4", Text: "vcl_pipe", IsCorrect: false, Explanation: "Used for pass-through, not typical fetching logic."},
		},
	},
	{
		ID:         "cache-h-q2",
		Topic:      "Cache and CDN",
		Difficulty: "hard",
		Text:       "Which HTTP/2 feature allows multiple responses to be sent for a single client request, improving CDN performance?",
		Options: []Option{
			{ID: "opt1", Text: "Server push", IsCorrect: true, Explanation: "Correct. Server push proactively sends resources."},
			{ID: "opt2", Text: "Header compression", IsCorrect: false, Explanation: "Compression saves bytes but not multiple responses."},
			{ID: "opt3", Text: "Stream prioritization", IsCorrect: false, Explanation: "Priorities reorder, not push extra responses."},
			{ID: "opt4", Text: "FIN flag", IsCorrect: false, Explanation: "TCP FIN closes connection; unrelated."},
		},
	},

	// ──────────────────────── Monitoring ────────────────────────
	{
		ID:         "mon-e-q1",
		Topic:      "Monitoring",
		Difficulty: "easy",
		Text:       "Prometheus stores data in a ___ database.",
		Options: []Option{
			{ID: "opt1", Text: "Time-series", IsCorrect: true, Explanation: "Correct. Prometheus is a TSDB."},
			{ID: "opt2", Text: "Graph", IsCorrect: false, Explanation: "OpenTSDB is closer but still TS."},
			{ID: "opt3", Text: "Relational", IsCorrect: false, Explanation: "TSDBs are not classic RDBMS."},
			{ID: "opt4", Text: "Document", IsCorrect: false, Explanation: "Document DBs store JSON-like docs."},
		},
	},
	{
		ID:         "mon-e-q2",
		Topic:      "Monitoring",
		Difficulty: "easy",
		Text:       "Which protocol do many network devices use to send traps to monitoring systems?",
		Options: []Option{
			{ID: "opt1", Text: "SNMP", IsCorrect: true, Explanation: "Correct. Simple Network Management Protocol traps notify events."},
			{ID: "opt2", Text: "SMTP", IsCorrect: false, Explanation: "SMTP sends email."},
			{ID: "opt3", Text: "SSH", IsCorrect: false, Explanation: "SSH is for remote shell."},
			{ID: "opt4", Text: "HTTP", IsCorrect: false, Explanation: "HTTP is web, not traditional trap."},
		},
	},
	{
		ID:         "mon-m-q1",
		Topic:      "Monitoring",
		Difficulty: "medium",
		Text:       "Grafana's alerting rule evaluates a query that must be in which state to fire an alert?",
		Options: []Option{
			{ID: "opt1", Text: "Alerting", IsCorrect: true, Explanation: "Correct. When evaluation returns 'alerting' threshold is breached."},
			{ID: "opt2", Text: "No data", IsCorrect: false, Explanation: "No data can optionally trigger alert but not default."},
			{ID: "opt3", Text: "OK", IsCorrect: false, Explanation: "OK means healthy."},
			{ID: "opt4", Text: "Paused", IsCorrect: false, Explanation: "Paused disables alerts."},
		},
	},
	{
		ID:         "mon-m-q2",
		Topic:      "Monitoring",
		Difficulty: "medium",
		Text:       "In PromQL, what does the 'rate()' function compute?",
		Options: []Option{
			{ID: "opt1", Text: "Per-second average increase of a counter", IsCorrect: true, Explanation: "Correct. rate calculates increase over time divided by seconds."},
			{ID: "opt2", Text: "Instantaneous gauge value", IsCorrect: false, Explanation: "That's direct selection of a gauge metric."},
			{ID: "opt3", Text: "Total sum of counter since start", IsCorrect: false, Explanation: "Use 'sum' over counter."},
			{ID: "opt4", Text: "Histogram quantile", IsCorrect: false, Explanation: "Use 'histogram_quantile' function."},
		},
	},
	{
		ID:         "mon-h-q1",
		Topic:      "Monitoring",
		Difficulty: "hard",
		Text:       "Which algorithm does Prometheus use to compress time-series chunks on disk?",
		Options: []Option{
			{ID: "opt1", Text: "Gorilla", IsCorrect: true, Explanation: "Correct. Facebook's Gorilla compression adapted for TSDB."},
			{ID: "opt2", Text: "LZMA", IsCorrect: false, Explanation: "Generic compression not used by Prometheus TSDB."},
			{ID: "opt3", Text: "Snappy", IsCorrect: false, Explanation: "Snappy compresses WAL but not TS chunks."},
			{ID: "opt4", Text: "Brotli", IsCorrect: false, Explanation: "Not used in Prometheus internal storage."},
		},
	},
	{
		ID:         "mon-h-q2",
		Topic:      "Monitoring",
		Difficulty: "hard",
		Text:       "In OpenTelemetry, spans are organized into ___ forming a complete trace.",
		Options: []Option{
			{ID: "opt1", Text: "Parent-child trees", IsCorrect: true, Explanation: "Correct. Spans nest to form a tree."},
			{ID: "opt2", Text: "Flat lists", IsCorrect: false, Explanation: "They have hierarchy, not flat."},
			{ID: "opt3", Text: "Ring buffers", IsCorrect: false, Explanation: "Storage implementation detail, not conceptual."},
			{ID: "opt4", Text: "Hash maps", IsCorrect: false, Explanation: "Spans can be indexed but topology is tree."},
		},
	},

	// ──────────────────────── Admin and Ops ────────────────────────
	{
		ID:         "ops-e-q1",
		Topic:      "Admin and Ops",
		Difficulty: "easy",
		Text:       "Which Linux command shows current disk usage of directories?",
		Options: []Option{
			{ID: "opt1", Text: "du", IsCorrect: true, Explanation: "Correct. 'du' summarizes disk usage."},
			{ID: "opt2", Text: "df", IsCorrect: false, Explanation: "'df' shows filesystem free space."},
			{ID: "opt3", Text: "ls", IsCorrect: false, Explanation: "Lists directory contents."},
			{ID: "opt4", Text: "top", IsCorrect: false, Explanation: "Shows running processes."},
		},
	},
	{
		ID:         "ops-e-q2",
		Topic:      "Admin and Ops",
		Difficulty: "easy",
		Text:       "Cron expressions schedule tasks on which operating systems by default?",
		Options: []Option{
			{ID: "opt1", Text: "Unix/Linux", IsCorrect: true, Explanation: "Correct. Cron is standard on Unix-like OSes."},
			{ID: "opt2", Text: "Windows only", IsCorrect: false, Explanation: "Task Scheduler is Windows equivalent."},
			{ID: "opt3", Text: "macOS exclusively", IsCorrect: false, Explanation: "macOS is Unix-based, but cron is not exclusive."},
			{ID: "opt4", Text: "Android only", IsCorrect: false, Explanation: "Android uses different schedulers."},
		},
	},
	{
		ID:         "ops-m-q1",
		Topic:      "Admin and Ops",
		Difficulty: "medium",
		Text:       "What does the 'systemctl isolate rescue.target' command achieve?",
		Options: []Option{
			{ID: "opt1", Text: "Brings the system into single-user rescue mode", IsCorrect: true, Explanation: "Correct. Only basic services with shell login."},
			{ID: "opt2", Text: "Reboots the system", IsCorrect: false, Explanation: "Use 'systemctl reboot' for that."},
			{ID: "opt3", Text: "Restarts networking only", IsCorrect: false, Explanation: "Use 'systemctl restart network' or similar."},
			{ID: "opt4", Text: "Shuts down immediately", IsCorrect: false, Explanation: "Use 'systemctl poweroff'."},
		},
	},
	{
		ID:         "ops-m-q2",
		Topic:      "Admin and Ops",
		Difficulty: "medium",
		Text:       "The tool 'htop' is an enhanced version of which classic Unix utility?",
		Options: []Option{
			{ID: "opt1", Text: "top", IsCorrect: true, Explanation: "Correct. htop adds colors, scrolling, and killing processes interactively."},
			{ID: "opt2", Text: "ps", IsCorrect: false, Explanation: "ps lists processes but isn't interactive like top."},
			{ID: "opt3", Text: "vmstat", IsCorrect: false, Explanation: "vmstat shows virtual memory stats."},
			{ID: "opt4", Text: "iostat", IsCorrect: false, Explanation: "iostat shows I/O stats."},
		},
	},
	{
		ID:         "ops-h-q1",
		Topic:      "Admin and Ops",
		Difficulty: "hard",
		Text:       "Which Linux kernel parameter toggles address space layout randomization (ASLR)?",
		Options: []Option{
			{ID: "opt1", Text: "/proc/sys/kernel/randomize_va_space", IsCorrect: true, Explanation: "Correct. Writing 0-2 configures ASLR modes."},
			{ID: "opt2", Text: "/proc/sys/kernel/panic", IsCorrect: false, Explanation: "Controls reboot after panic."},
			{ID: "opt3", Text: "/proc/sys/vm/swappiness", IsCorrect: false, Explanation: "Adjusts swapping aggressiveness."},
			{ID: "opt4", Text: "/proc/sys/net/ipv4/ip_forward", IsCorrect: false, Explanation: "Enables IP forwarding."},
		},
	},
	{
		ID:         "ops-h-q2",
		Topic:      "Admin and Ops",
		Difficulty: "hard",
		Text:       "Which log level numeric value represents 'critical' in syslog?",
		Options: []Option{
			{ID: "opt1", Text: "2", IsCorrect: true, Explanation: "Correct. Syslog levels: 0 emerg, 1 alert, 2 crit."},
			{ID: "opt2", Text: "4", IsCorrect: false, Explanation: "4 is warning."},
			{ID: "opt3", Text: "3", IsCorrect: false, Explanation: "3 is error."},
			{ID: "opt4", Text: "5", IsCorrect: false, Explanation: "5 is notice."},
		},
	},

	// ──────────────────────── Kubernetes ────────────────────────
	{
		ID:         "k8s-e-q1",
		Topic:      "Kubernetes",
		Difficulty: "easy",
		Text:       "Which Kubernetes object ensures the desired number of pod replicas?",
		Options: []Option{
			{ID: "opt1", Text: "Deployment", IsCorrect: true, Explanation: "Correct. Deployment manages ReplicaSets."},
			{ID: "opt2", Text: "Service", IsCorrect: false, Explanation: "Service exposes pods but doesn't scale them."},
			{ID: "opt3", Text: "ConfigMap", IsCorrect: false, Explanation: "Stores configuration data."},
			{ID: "opt4", Text: "Ingress", IsCorrect: false, Explanation: "Manages external HTTP(S) access."},
		},
	},
	{
		ID:         "k8s-e-q2",
		Topic:      "Kubernetes",
		Difficulty: "easy",
		Text:       "Which command-line tool interacts with a Kubernetes API server?",
		Options: []Option{
			{ID: "opt1", Text: "kubectl", IsCorrect: true, Explanation: "Correct. kubectl manages cluster resources."},
			{ID: "opt2", Text: "docker", IsCorrect: false, Explanation: "Docker manages containers, not K8s objects directly."},
			{ID: "opt3", Text: "helm", IsCorrect: false, Explanation: "Helm is a package manager on top of K8s."},
			{ID: "opt4", Text: "etcdctl", IsCorrect: false, Explanation: "etcdctl interacts with etcd key-value store."},
		},
	},
	{
		ID:         "k8s-m-q1",
		Topic:      "Kubernetes",
		Difficulty: "medium",
		Text:       "Which controller monitors Node health and replaces failed Nodes in managed services like GKE and EKS?",
		Options: []Option{
			{ID: "opt1", Text: "Cluster Autoscaler", IsCorrect: true, Explanation: "Correct. It deletes unhealthy nodes and adds new ones."},
			{ID: "opt2", Text: "Horizontal Pod Autoscaler", IsCorrect: false, Explanation: "HPA scales pods, not nodes."},
			{ID: "opt3", Text: "Vertical Pod Autoscaler", IsCorrect: false, Explanation: "VPA adjusts pod resource requests."},
			{ID: "opt4", Text: "ReplicaSet controller", IsCorrect: false, Explanation: "ReplicaSet manages pods, not nodes."},
		},
	},
	{
		ID:         "k8s-m-q2",
		Topic:      "Kubernetes",
		Difficulty: "medium",
		Text:       "What field in a pod spec defines the container image pull policy?",
		Options: []Option{
			{ID: "opt1", Text: "imagePullPolicy", IsCorrect: true, Explanation: "Correct. Options: Always, IfNotPresent, Never."},
			{ID: "opt2", Text: "restartPolicy", IsCorrect: false, Explanation: "Controls pod restart semantics."},
			{ID: "opt3", Text: "command", IsCorrect: false, Explanation: "Overrides container ENTRYPOINT."},
			{ID: "opt4", Text: "nodeSelector", IsCorrect: false, Explanation: "Constrains scheduling to nodes."},
		},
	},
	{
		ID:         "k8s-h-q1",
		Topic:      "Kubernetes",
		Difficulty: "hard",
		Text:       "Etcd, used as Kubernetes' backing store, uses which consensus algorithm?",
		Options: []Option{
			{ID: "opt1", Text: "Raft", IsCorrect: true, Explanation: "Correct. Raft ensures strong consistency."},
			{ID: "opt2", Text: "Paxos", IsCorrect: false, Explanation: "Raft is easier to understand and is used in etcd."},
			{ID: "opt3", Text: "Gossip", IsCorrect: false, Explanation: "Gossip protocols offer eventual consistency, not consensus."},
			{ID: "opt4", Text: "Two-phase commit", IsCorrect: false, Explanation: "2PC coordinates transactions, not leader election."},
		},
	},
	{
		ID:         "k8s-h-q2",
		Topic:      "Kubernetes",
		Difficulty: "hard",
		Text:       "Which admission controller verifies image signatures before pods are scheduled?",
		Options: []Option{
			{ID: "opt1", Text: "ImagePolicyWebhook", IsCorrect: true, Explanation: "Correct. It calls an external service to validate images."},
			{ID: "opt2", Text: "PodSecurityPolicy", IsCorrect: false, Explanation: "PSP restricts pod permissions."},
			{ID: "opt3", Text: "ResourceQuota", IsCorrect: false, Explanation: "ResourceQuota limits resource consumption."},
			{ID: "opt4", Text: "NodeRestriction", IsCorrect: false, Explanation: "NodeRestriction limits kubelet API actions."},
		},
	},

	// ──────────────────────── Linux ────────────────────────
	{
		ID:         "lin-e-q1",
		Topic:      "Linux",
		Difficulty: "easy",
		Text:       "Which command prints the current directory path?",
		Options: []Option{
			{ID: "opt1", Text: "pwd", IsCorrect: true, Explanation: "Correct. 'pwd' prints working directory."},
			{ID: "opt2", Text: "cd", IsCorrect: false, Explanation: "cd changes directories."},
			{ID: "opt3", Text: "ls", IsCorrect: false, Explanation: "Lists files."},
			{ID: "opt4", Text: "mkdir", IsCorrect: false, Explanation: "Creates directories."},
		},
	},
	{
		ID:         "lin-e-q2",
		Topic:      "Linux",
		Difficulty: "easy",
		Text:       "In 'chmod 755 file', what permissions does '7' grant the owner?",
		Options: []Option{
			{ID: "opt1", Text: "Read, write, and execute", IsCorrect: true, Explanation: "Correct. 7 = 4+2+1."},
			{ID: "opt2", Text: "Read and execute only", IsCorrect: false, Explanation: "That's 5."},
			{ID: "opt3", Text: "Write only", IsCorrect: false, Explanation: "That's 2."},
			{ID: "opt4", Text: "Execute only", IsCorrect: false, Explanation: "That's 1."},
		},
	},
	{
		ID:         "lin-m-q1",
		Topic:      "Linux",
		Difficulty: "medium",
		Text:       "Which file lists mounted filesystems at boot and is used by the mount command when no options are given?",
		Options: []Option{
			{ID: "opt1", Text: "/etc/fstab", IsCorrect: true, Explanation: "Correct. fstab defines static mounts."},
			{ID: "opt2", Text: "/etc/mtab", IsCorrect: false, Explanation: "mtab shows currently mounted FS."},
			{ID: "opt3", Text: "/proc/mounts", IsCorrect: false, Explanation: "proc lists current mounts too but not static config."},
			{ID: "opt4", Text: "/etc/filesystems", IsCorrect: false, Explanation: "Lists filesystem types allowed."},
		},
	},
	{
		ID:         "lin-m-q2",
		Topic:      "Linux",
		Difficulty: "medium",
		Text:       "Which signal number is associated with 'kill -9'?",
		Options: []Option{
			{ID: "opt1", Text: "SIGKILL (9)", IsCorrect: true, Explanation: "Correct. Cannot be caught or ignored."},
			{ID: "opt2", Text: "SIGTERM (15)", IsCorrect: false, Explanation: "Default termination signal."},
			{ID: "opt3", Text: "SIGINT (2)", IsCorrect: false, Explanation: "Interrupt from keyboard."},
			{ID: "opt4", Text: "SIGHUP (1)", IsCorrect: false, Explanation: "Hangup detected."},
		},
	},
	{
		ID:         "lin-h-q1",
		Topic:      "Linux",
		Difficulty: "hard",
		Text:       "Which scheduling policy has the numeric value '6' in the /proc/<pid>/sched file?",
		Options: []Option{
			{ID: "opt1", Text: "SCHED_DEADLINE", IsCorrect: true, Explanation: "Correct. Policies: 0 normal, 1 FIFO, 2 RR, 3 batch, 5 idle, 6 DL."},
			{ID: "opt2", Text: "SCHED_FIFO", IsCorrect: false, Explanation: "FIFO = 1."},
			{ID: "opt3", Text: "SCHED_RR", IsCorrect: false, Explanation: "RR = 2."},
			{ID: "opt4", Text: "SCHED_BATCH", IsCorrect: false, Explanation: "Batch = 3."},
		},
	},
	{
		ID:         "lin-h-q2",
		Topic:      "Linux",
		Difficulty: "hard",
		Text:       "Which bpftrace probe type attaches to kernel function entry points?",
		Options: []Option{
			{ID: "opt1", Text: "kprobe", IsCorrect: true, Explanation: "Correct. kprobes fire on kernel functions."},
			{ID: "opt2", Text: "uprobes", IsCorrect: false, Explanation: "uprobes are for user-space functions."},
			{ID: "opt3", Text: "tracepoint", IsCorrect: false, Explanation: "Tracepoints are static hooks."},
			{ID: "opt4", Text: "usdt", IsCorrect: false, Explanation: "USDT probes instrument user statically-defined tracing."},
		},
	},

	// ──────────────────────── Pipelines & CI/CD ────────────────────────
	{
		ID:         "cicd-e-q1",
		Topic:      "Pipelines and CI/CD",
		Difficulty: "easy",
		Text:       "Which open-source tool by HashiCorp automates infrastructure provisioning and is often used in pipelines?",
		Options: []Option{
			{ID: "opt1", Text: "Terraform", IsCorrect: true, Explanation: "Correct. It's IaC in many CI/CD flows."},
			{ID: "opt2", Text: "Chef", IsCorrect: false, Explanation: "Chef is configuration management, not HashiCorp."},
			{ID: "opt3", Text: "Puppet", IsCorrect: false, Explanation: "Also config management."},
			{ID: "opt4", Text: "Maven", IsCorrect: false, Explanation: "Maven builds Java projects."},
		},
	},
	{
		ID:         "cicd-e-q2",
		Topic:      "Pipelines and CI/CD",
		Difficulty: "easy",
		Text:       "In GitHub Actions, a workflow file is written in which format?",
		Options: []Option{
			{ID: "opt1", Text: "YAML", IsCorrect: true, Explanation: "Correct. Workflows reside under .github/workflows/*.yml."},
			{ID: "opt2", Text: "XML", IsCorrect: false, Explanation: "Maven POMs use XML."},
			{ID: "opt3", Text: "JSON", IsCorrect: false, Explanation: "JSON isn't used for GHA workflows."},
			{ID: "opt4", Text: "INI", IsCorrect: false, Explanation: "INI rarely used here."},
		},
	},
	{
		ID:         "cicd-m-q1",
		Topic:      "Pipelines and CI/CD",
		Difficulty: "medium",
		Text:       "Which Jenkins plugin allows defining pipelines in a 'Jenkinsfile' using Groovy DSL?",
		Options: []Option{
			{ID: "opt1", Text: "Pipeline (Workflow) plugin", IsCorrect: true, Explanation: "Correct. Enables scripted and declarative pipelines."},
			{ID: "opt2", Text: "Blue Ocean", IsCorrect: false, Explanation: "Blue Ocean is a UI for pipelines."},
			{ID: "opt3", Text: "Job DSL", IsCorrect: false, Explanation: "Job DSL creates jobs programmatically, not Jenkinsfile."},
			{ID: "opt4", Text: "AnsiColor", IsCorrect: false, Explanation: "Adds color to console output."},
		},
	},
	{
		ID:         "cicd-m-q2",
		Topic:      "Pipelines and CI/CD",
		Difficulty: "medium",
		Text:       "Which GitLab CI keyword specifies a Docker image to use for running a job?",
		Options: []Option{
			{ID: "opt1", Text: "image", IsCorrect: true, Explanation: "Correct. Placed at job or global level."},
			{ID: "opt2", Text: "stage", IsCorrect: false, Explanation: "Defines pipeline stage order."},
			{ID: "opt3", Text: "script", IsCorrect: false, Explanation: "Script lists shell commands."},
			{ID: "opt4", Text: "artifacts", IsCorrect: false, Explanation: "Specifies files to save after job."},
		},
	},
	{
		ID:         "cicd-h-q1",
		Topic:      "Pipelines and CI/CD",
		Difficulty: "hard",
		Text:       "Spinnaker's 'bake' stage primarily integrates with which tool to build VM images?",
		Options: []Option{
			{ID: "opt1", Text: "HashiCorp Packer", IsCorrect: true, Explanation: "Correct. Packer templates build AMIs or other images."},
			{ID: "opt2", Text: "Docker Buildx", IsCorrect: false, Explanation: "Buildx builds container images, not VM images."},
			{ID: "opt3", Text: "Ansible-playbook", IsCorrect: false, Explanation: "Ansible configures servers, not build base images."},
			{ID: "opt4", Text: "Chef InSpec", IsCorrect: false, Explanation: "InSpec tests infrastructure, not builds images."},
		},
	},
	{
		ID:         "cicd-h-q2",
		Topic:      "Pipelines and CI/CD",
		Difficulty: "hard",
		Text:       "Which concept in Argo CD ensures the actual state matches the desired state declared in Git?",
		Options: []Option{
			{ID: "opt1", Text: "Declarative GitOps reconciliation loop", IsCorrect: true, Explanation: "Correct. Argo continuously reconciles state."},
			{ID: "opt2", Text: "Blue-green deployment", IsCorrect: false, Explanation: "Blue-green is a release strategy."},
			{ID: "opt3", Text: "Canary analysis", IsCorrect: false, Explanation: "Canary is progressive delivery."},
			{ID: "opt4", Text: "Server-side apply", IsCorrect: false, Explanation: "K8s SSA is differencing mechanism; not full reconciliation loop."},
		},
	},

	// ──────────────────────── APIs ────────────────────────
	{
		ID:         "api-e-q1",
		Topic:      "APIs",
		Difficulty: "easy",
		Text:       "Which HTTP method is typically used to create a new resource?",
		Options: []Option{
			{ID: "opt1", Text: "POST", IsCorrect: true, Explanation: "Correct. POST creates or submits data."},
			{ID: "opt2", Text: "GET", IsCorrect: false, Explanation: "GET retrieves data without side effects."},
			{ID: "opt3", Text: "PUT", IsCorrect: false, Explanation: "PUT usually replaces a resource."},
			{ID: "opt4", Text: "DELETE", IsCorrect: false, Explanation: "DELETE removes a resource."},
		},
	},
	{
		ID:         "api-e-q2",
		Topic:      "APIs",
		Difficulty: "easy",
		Text:       "OpenAPI (Swagger) specifications are commonly written in which two human-readable formats?",
		Options: []Option{
			{ID: "opt1", Text: "YAML and JSON", IsCorrect: true, Explanation: "Correct. Both describe schemas."},
			{ID: "opt2", Text: "XML and TOML", IsCorrect: false, Explanation: "Possible but not common for Swagger."},
			{ID: "opt3", Text: "Markdown only", IsCorrect: false, Explanation: "Markdown is documentation, not spec serialization."},
			{ID: "opt4", Text: "Binary Protobufs", IsCorrect: false, Explanation: "Protobufs are for gRPC, not OpenAPI."},
		},
	},
	{
		ID:         "api-m-q1",
		Topic:      "APIs",
		Difficulty: "medium",
		Text:       "Which HTTP status code indicates 'Unprocessable Entity' often returned by validation errors?",
		Options: []Option{
			{ID: "opt1", Text: "422", IsCorrect: true, Explanation: "Correct. RFC 4918."},
			{ID: "opt2", Text: "409", IsCorrect: false, Explanation: "Conflict."},
			{ID: "opt3", Text: "400", IsCorrect: false, Explanation: "Bad Request generic."},
			{ID: "opt4", Text: "403", IsCorrect: false, Explanation: "Forbidden."},
		},
	},
	{
		ID:         "api-m-q2",
		Topic:      "APIs",
		Difficulty: "medium",
		Text:       "In gRPC, what is the default serialization format for messages?",
		Options: []Option{
			{ID: "opt1", Text: "Protocol Buffers", IsCorrect: true, Explanation: "Correct. Compact binary schema."},
			{ID: "opt2", Text: "JSON", IsCorrect: false, Explanation: "JSON is for REST APIs, gRPC uses protobuf."},
			{ID: "opt3", Text: "XML", IsCorrect: false, Explanation: "XML SOAP is legacy."},
			{ID: "opt4", Text: "YAML", IsCorrect: false, Explanation: "Not default format."},
		},
	},
	{
		ID:         "api-h-q1",
		Topic:      "APIs",
		Difficulty: "hard",
		Text:       "Which OAuth 2.0 grant type is recommended for single-page applications without a backend?",
		Options: []Option{
			{ID: "opt1", Text: "Authorization Code with PKCE", IsCorrect: true, Explanation: "Correct. PKCE secures public clients."},
			{ID: "opt2", Text: "Implicit grant", IsCorrect: false, Explanation: "Implicit is discouraged due to token leakage risk."},
			{ID: "opt3", Text: "Client credentials", IsCorrect: false, Explanation: "Used for machine-to-machine."},
			{ID: "opt4", Text: "Resource owner password", IsCorrect: false, Explanation: "Deprecated and insecure for SPAs."},
		},
	},
	{
		ID:         "api-h-q2",
		Topic:      "APIs",
		Difficulty: "hard",
		Text:       "HATEOAS is a constraint of which software architectural style?",
		Options: []Option{
			{ID: "opt1", Text: "REST", IsCorrect: true, Explanation: "Correct. Hypermedia as the Engine of Application State."},
			{ID: "opt2", Text: "GraphQL", IsCorrect: false, Explanation: "GraphQL is a query language, not REST property."},
			{ID: "opt3", Text: "SOAP", IsCorrect: false, Explanation: "SOAP is XML-based protocol."},
			{ID: "opt4", Text: "gRPC", IsCorrect: false, Explanation: "gRPC uses HTTP/2 and protobufs."},
		},
	},

	// ──────────────────────── Terraform ────────────────────────
	{
		ID:         "tf-e-q1",
		Topic:      "Terraform",
		Difficulty: "easy",
		Text:       "Terraform configuration files use which filename extension?",
		Options: []Option{
			{ID: "opt1", Text: ".tf", IsCorrect: true, Explanation: "Correct. Modules may also use .tf.json."},
			{ID: "opt2", Text: ".yaml", IsCorrect: false, Explanation: "Used by Ansible, Kubernetes, etc."},
			{ID: "opt3", Text: ".hcl", IsCorrect: false, Explanation: "HCL is the language, but Terraform files are .tf."},
			{ID: "opt4", Text: ".ini", IsCorrect: false, Explanation: "INI is unrelated."},
		},
	},
	{
		ID:         "tf-e-q2",
		Topic:      "Terraform",
		Difficulty: "easy",
		Text:       "Which command initializes a Terraform working directory?",
		Options: []Option{
			{ID: "opt1", Text: "terraform init", IsCorrect: true, Explanation: "Correct. Downloads providers, modules, etc."},
			{ID: "opt2", Text: "terraform plan", IsCorrect: false, Explanation: "Shows execution plan but needs init first."},
			{ID: "opt3", Text: "terraform apply", IsCorrect: false, Explanation: "Applies changes."},
			{ID: "opt4", Text: "terraform destroy", IsCorrect: false, Explanation: "Destroys managed infrastructure."},
		},
	},
	{
		ID:         "tf-m-q1",
		Topic:      "Terraform",
		Difficulty: "medium",
		Text:       "A 'data' block in Terraform is used for what purpose?",
		Options: []Option{
			{ID: "opt1", Text: "Reading existing resources that were not created by Terraform", IsCorrect: true, Explanation: "Correct. Data sources query external objects."},
			{ID: "opt2", Text: "Creating new resources", IsCorrect: false, Explanation: "Use 'resource' blocks to create."},
			{ID: "opt3", Text: "Defining variables", IsCorrect: false, Explanation: "Variables are declared with 'variable'."},
			{ID: "opt4", Text: "Outputting values", IsCorrect: false, Explanation: "Use 'output' blocks."},
		},
	},
	{
		ID:         "tf-m-q2",
		Topic:      "Terraform",
		Difficulty: "medium",
		Text:       "What does the 'terraform taint' command mark a resource for?",
		Options: []Option{
			{ID: "opt1", Text: "Forced recreation on next apply", IsCorrect: true, Explanation: "Correct. Tainted resources will be destroyed and recreated."},
			{ID: "opt2", Text: "Deletion without replacement", IsCorrect: false, Explanation: "Destroy does that."},
			{ID: "opt3", Text: "Import into the state file", IsCorrect: false, Explanation: "Use 'terraform import'."},
			{ID: "opt4", Text: "Locking state", IsCorrect: false, Explanation: "State locking handled automatically."},
		},
	},
	{
		ID:         "tf-h-q1",
		Topic:      "Terraform",
		Difficulty: "hard",
		Text:       "Which backend enables Terraform state storage with state locking using DynamoDB?",
		Options: []Option{
			{ID: "opt1", Text: "S3 backend with DynamoDB lock table", IsCorrect: true, Explanation: "Correct. S3 stores state; DynamoDB coordinates locks."},
			{ID: "opt2", Text: "Local backend", IsCorrect: false, Explanation: "Local has no locking."},
			{ID: "opt3", Text: "Consul backend", IsCorrect: false, Explanation: "Consul locks via sessions, not DynamoDB."},
			{ID: "opt4", Text: "GCS backend", IsCorrect: false, Explanation: "GCS uses Cloud Storage with optional locks via Cloud Firestore workaround."},
		},
	},
	{
		ID:         "tf-h-q2",
		Topic:      "Terraform",
		Difficulty: "hard",
		Text:       "In Terraform 1.6+, what feature lets providers declare dependencies on resources managed by other providers?",
		Options: []Option{
			{ID: "opt1", Text: "Provider-scoped data source 'externals'", IsCorrect: true, Explanation: "Correct. Cross-provider data dependencies supported via externals."},
			{ID: "opt2", Text: "Variable interpolation", IsCorrect: false, Explanation: "Interpolation existed earlier and isn't provider-level dependency."},
			{ID: "opt3", Text: "Module outputs", IsCorrect: false, Explanation: "Modules help but not provider dependency declaration."},
			{ID: "opt4", Text: "Backend override files", IsCorrect: false, Explanation: "Overrides backend but not provider dependencies."},
		},
	},

	// ──────────────────────── Ansible ────────────────────────
	{
		ID:         "ans-e-q1",
		Topic:      "Ansible",
		Difficulty: "easy",
		Text:       "Ansible playbooks are written in which language?",
		Options: []Option{
			{ID: "opt1", Text: "YAML", IsCorrect: true, Explanation: "Correct. YAML declaratively defines tasks."},
			{ID: "opt2", Text: "Python", IsCorrect: false, Explanation: "Modules are often Python, but playbooks are YAML."},
			{ID: "opt3", Text: "JSON", IsCorrect: false, Explanation: "JSON can be used but uncommon."},
			{ID: "opt4", Text: "INI", IsCorrect: false, Explanation: "INI is for inventory files, not playbooks."},
		},
	},
	{
		ID:         "ans-e-q2",
		Topic:      "Ansible",
		Difficulty: "easy",
		Text:       "Which command executes an ad-hoc Ansible task on hosts?",
		Options: []Option{
			{ID: "opt1", Text: "ansible", IsCorrect: true, Explanation: "Correct. 'ansible' runs a single module ad-hoc."},
			{ID: "opt2", Text: "ansible-playbook", IsCorrect: false, Explanation: "Runs full playbooks."},
			{ID: "opt3", Text: "ansible-vault", IsCorrect: false, Explanation: "Manages encrypted secrets."},
			{ID: "opt4", Text: "ansible-galaxy", IsCorrect: false, Explanation: "Downloads roles."},
		},
	},
	{
		ID:         "ans-m-q1",
		Topic:      "Ansible",
		Difficulty: "medium",
		Text:       "What is the default location of the Ansible inventory file?",
		Options: []Option{
			{ID: "opt1", Text: "/etc/ansible/hosts", IsCorrect: true, Explanation: "Correct. Inventory defines host groups."},
			{ID: "opt2", Text: "/etc/ansible/ansible.cfg", IsCorrect: false, Explanation: "Config file path."},
			{ID: "opt3", Text: "~/.ansible/inventory", IsCorrect: false, Explanation: "Not default."},
			{ID: "opt4", Text: "/usr/local/ansible/hosts", IsCorrect: false, Explanation: "Not default path."},
		},
	},
	{
		ID:         "ans-m-q2",
		Topic:      "Ansible",
		Difficulty: "medium",
		Text:       "Which Ansible feature encrypts sensitive variables and files at rest?",
		Options: []Option{
			{ID: "opt1", Text: "Ansible Vault", IsCorrect: true, Explanation: "Correct. Vault encrypts data with AES."},
			{ID: "opt2", Text: "Ansible Galaxy", IsCorrect: false, Explanation: "Galaxy is a role sharing site."},
			{ID: "opt3", Text: "Ansible Collections", IsCorrect: false, Explanation: "Collections bundle roles/modules."},
			{ID: "opt4", Text: "Ansible Facts", IsCorrect: false, Explanation: "Facts gather system info."},
		},
	},
	{
		ID:         "ans-h-q1",
		Topic:      "Ansible",
		Difficulty: "hard",
		Text:       "What does the 'strategy: free' option achieve in a playbook?",
		Options: []Option{
			{ID: "opt1", Text: "Allows tasks to run asynchronously on hosts without waiting for others", IsCorrect: true, Explanation: "Correct. Hosts don't block on each other."},
			{ID: "opt2", Text: "Runs tasks in serial one host at a time", IsCorrect: false, Explanation: "That's 'serial'."},
			{ID: "opt3", Text: "Locks inventory to prevent parallelism", IsCorrect: false, Explanation: "Opposite effect."},
			{ID: "opt4", Text: "Randomizes task order", IsCorrect: false, Explanation: "Tasks still in order, just not synchronized across hosts."},
		},
	},
	{
		ID:         "ans-h-q2",
		Topic:      "Ansible",
		Difficulty: "hard",
		Text:       "When writing a custom dynamic inventory script, which environment variable provides the path to a temporary directory for cache?",
		Options: []Option{
			{ID: "opt1", Text: "ANSIBLE_TMPDIR", IsCorrect: true, Explanation: "Correct. Script can cache output there."},
			{ID: "opt2", Text: "ANSIBLE_CONFIG", IsCorrect: false, Explanation: "Points to config file."},
			{ID: "opt3", Text: "ANSIBLE_VAULT_PASSWORD_FILE", IsCorrect: false, Explanation: "Vault password file path."},
			{ID: "opt4", Text: "ANSIBLE_STDOUT_CALLBACK", IsCorrect: false, Explanation: "Specifies stdout callback plugin."},
		},
	},

	// ──────────────────────── Azure ────────────────────────
	{
		ID:         "az-e-q1",
		Topic:      "Azure",
		Difficulty: "easy",
		Text:       "Azure virtual machines are billed per ___.",
		Options: []Option{
			{ID: "opt1", Text: "Second (for Linux VMs)", IsCorrect: true, Explanation: "Correct. Azure bills per second with a one-minute minimum."},
			{ID: "opt2", Text: "Month upfront only", IsCorrect: false, Explanation: "Pay-as-you-go is per second."},
			{ID: "opt3", Text: "Packet transferred", IsCorrect: false, Explanation: "Bandwidth billed separately."},
			{ID: "opt4", Text: "User account", IsCorrect: false, Explanation: "Account not a billing metric."},
		},
	},
	{
		ID:         "az-e-q2",
		Topic:      "Azure",
		Difficulty: "easy",
		Text:       "Which service stores unstructured object data similar to AWS S3?",
		Options: []Option{
			{ID: "opt1", Text: "Azure Blob Storage", IsCorrect: true, Explanation: "Correct. Blob storage provides buckets called containers."},
			{ID: "opt2", Text: "Azure SQL Database", IsCorrect: false, Explanation: "SQL is relational."},
			{ID: "opt3", Text: "Azure Functions", IsCorrect: false, Explanation: "Functions is serverless compute."},
			{ID: "opt4", Text: "Azure DevOps", IsCorrect: false, Explanation: "DevOps is a suite of services."},
		},
	},
	{
		ID:         "az-m-q1",
		Topic:      "Azure",
		Difficulty: "medium",
		Text:       "What is the replication option of an Azure Storage account that provides 3 synchronous copies within a single region?",
		Options: []Option{
			{ID: "opt1", Text: "Locally-redundant storage (LRS)", IsCorrect: true, Explanation: "Correct. Three copies in one DC."},
			{ID: "opt2", Text: "Zone-redundant storage (ZRS)", IsCorrect: false, Explanation: "Copies across zones."},
			{ID: "opt3", Text: "Geo-redundant storage (GRS)", IsCorrect: false, Explanation: "Replicates cross-region."},
			{ID: "opt4", Text: "Read-access GRS", IsCorrect: false, Explanation: "Adds secondary read endpoint."},
		},
	},
	{
		ID:         "az-m-q2",
		Topic:      "Azure",
		Difficulty: "medium",
		Text:       "Which Azure service provides managed Kubernetes clusters?",
		Options: []Option{
			{ID: "opt1", Text: "AKS (Azure Kubernetes Service)", IsCorrect: true, Explanation: "Correct. Managed control plane."},
			{ID: "opt2", Text: "App Service", IsCorrect: false, Explanation: "App Service hosts web apps, not K8s."},
			{ID: "opt3", Text: "Service Fabric", IsCorrect: false, Explanation: "Microservices platform, not K8s."},
			{ID: "opt4", Text: "Container Instances", IsCorrect: false, Explanation: "Runs single containers, not clusters."},
		},
	},
	{
		ID:         "az-h-q1",
		Topic:      "Azure",
		Difficulty: "hard",
		Text:       "In Azure AD, what protocol underlies Conditional Access sign-in risk evaluation using tokens?",
		Options: []Option{
			{ID: "opt1", Text: "OAuth 2.0", IsCorrect: true, Explanation: "Correct. CA policies assess OAuth tokens."},
			{ID: "opt2", Text: "RADIUS", IsCorrect: false, Explanation: "Legacy auth for networking."},
			{ID: "opt3", Text: "Kerberos", IsCorrect: false, Explanation: "On-prem AD uses Kerberos."},
			{ID: "opt4", Text: "SAML 1.1", IsCorrect: false, Explanation: "SAML can integrate but CA risk uses OAuth tokens."},
		},
	},
	{
		ID:         "az-h-q2",
		Topic:      "Azure",
		Difficulty: "hard",
		Text:       "Which Terraform AzureRM resource creates a private DNS zone?",
		Options: []Option{
			{ID: "opt1", Text: "azurerm_private_dns_zone", IsCorrect: true, Explanation: "Correct. Resource manages private DNS."},
			{ID: "opt2", Text: "azurerm_dns_zone", IsCorrect: false, Explanation: "Public DNS zone."},
			{ID: "opt3", Text: "azurerm_network_dns_zone", IsCorrect: false, Explanation: "Non-existent."},
			{ID: "opt4", Text: "azurerm_private_zone", IsCorrect: false, Explanation: "Non-existent resource name."},
		},
	},

	// ──────────────────────── AWS ────────────────────────
	{
		ID:         "aws-e-q1",
		Topic:      "AWS",
		Difficulty: "easy",
		Text:       "What does EC2 stand for?",
		Options: []Option{
			{ID: "opt1", Text: "Elastic Compute Cloud", IsCorrect: true, Explanation: "Correct. Core virtual machine service."},
			{ID: "opt2", Text: "Edge Content Cache", IsCorrect: false, Explanation: "Not EC2."},
			{ID: "opt3", Text: "Elastic Container Cluster", IsCorrect: false, Explanation: "ECS/EKS manage containers."},
			{ID: "opt4", Text: "Enhanced Cloud Console", IsCorrect: false, Explanation: "Not a service."},
		},
	},
	{
		ID:         "aws-e-q2",
		Topic:      "AWS",
		Difficulty: "easy",
		Text:       "Which AWS service provides a managed MySQL-compatible database with horizontal scalability?",
		Options: []Option{
			{ID: "opt1", Text: "Amazon Aurora", IsCorrect: true, Explanation: "Correct. Aurora MySQL and PostgreSQL variants."},
			{ID: "opt2", Text: "DynamoDB", IsCorrect: false, Explanation: "DynamoDB is NoSQL key-value."},
			{ID: "opt3", Text: "Redshift", IsCorrect: false, Explanation: "Redshift is data warehouse."},
			{ID: "opt4", Text: "RDS for Oracle", IsCorrect: false, Explanation: "Not MySQL compatible."},
		},
	},
	{
		ID:         "aws-m-q1",
		Topic:      "AWS",
		Difficulty: "medium",
		Text:       "Which IAM entity is designed for temporary credentials and cannot directly own long-term access keys?",
		Options: []Option{
			{ID: "opt1", Text: "IAM role", IsCorrect: true, Explanation: "Correct. Roles are assumed, not logged into."},
			{ID: "opt2", Text: "IAM user", IsCorrect: false, Explanation: "Users have permanent credentials."},
			{ID: "opt3", Text: "IAM group", IsCorrect: false, Explanation: "Groups bundle permissions, no keys."},
			{ID: "opt4", Text: "Service control policy", IsCorrect: false, Explanation: "SCP applies to AWS organizations."},
		},
	},
	{
		ID:         "aws-m-q2",
		Topic:      "AWS",
		Difficulty: "medium",
		Text:       "S3 guarantees what durability percentage for objects?",
		Options: []Option{
			{ID: "opt1", Text: "99.999999999% (11 nines)", IsCorrect: true, Explanation: "Correct. Across multiple AZs."},
			{ID: "opt2", Text: "99.9%", IsCorrect: false, Explanation: "That's availability for some services."},
			{ID: "opt3", Text: "100%", IsCorrect: false, Explanation: "No service guarantees 100%."},
			{ID: "opt4", Text: "95%", IsCorrect: false, Explanation: "Durability is far higher."},
		},
	},
	{
		ID:         "aws-h-q1",
		Topic:      "AWS",
		Difficulty: "hard",
		Text:       "Which AWS service enables DNS-based traffic routing and health checks?",
		Options: []Option{
			{ID: "opt1", Text: "Route 53", IsCorrect: true, Explanation: "Correct. It is AWS's authoritative DNS."},
			{ID: "opt2", Text: "Global Accelerator", IsCorrect: false, Explanation: "Accelerator uses Anycast but not authoritative DNS routing policies."},
			{ID: "opt3", Text: "CloudFront", IsCorrect: false, Explanation: "CDN distribution requires DNS but not primary."},
			{ID: "opt4", Text: "Direct Connect", IsCorrect: false, Explanation: "Private networking service."},
		},
	},
	{
		ID:         "aws-h-q2",
		Topic:      "AWS",
		Difficulty: "hard",
		Text:       "What database engine does AWS Timestream target?",
		Options: []Option{
			{ID: "opt1", Text: "Time-series workloads", IsCorrect: true, Explanation: "Correct. Purpose-built TSDB."},
			{ID: "opt2", Text: "Graph queries", IsCorrect: false, Explanation: "Graph workloads use Neptune."},
			{ID: "opt3", Text: "Document storage", IsCorrect: false, Explanation: "Document DB uses DocumentDB."},
			{ID: "opt4", Text: "Data warehousing", IsCorrect: false, Explanation: "Redshift is data warehouse."},
		},
	},

	// ──────────────────────── GCP ────────────────────────
	{
		ID:         "gcp-e-q1",
		Topic:      "GCP",
		Difficulty: "easy",
		Text:       "Google Cloud virtual machines are provided by which service?",
		Options: []Option{
			{ID: "opt1", Text: "Compute Engine", IsCorrect: true, Explanation: "Correct. Compute Engine offers VMs."},
			{ID: "opt2", Text: "App Engine", IsCorrect: false, Explanation: "App Engine is PaaS."},
			{ID: "opt3", Text: "Cloud Functions", IsCorrect: false, Explanation: "Functions is serverless."},
			{ID: "opt4", Text: "Cloud Run", IsCorrect: false, Explanation: "Cloud Run runs containers."},
		},
	},
	{
		ID:         "gcp-e-q2",
		Topic:      "GCP",
		Difficulty: "easy",
		Text:       "Which GCP service is a managed, scalable, NoSQL document database?",
		Options: []Option{
			{ID: "opt1", Text: "Cloud Firestore", IsCorrect: true, Explanation: "Correct. Document database with realtime features."},
			{ID: "opt2", Text: "Cloud Spanner", IsCorrect: false, Explanation: "Spanner is relational, global scale."},
			{ID: "opt3", Text: "Bigtable", IsCorrect: false, Explanation: "Bigtable is wide-column."},
			{ID: "opt4", Text: "Cloud SQL", IsCorrect: false, Explanation: "Managed MySQL/Postgres."},
		},
	},
	{
		ID:         "gcp-m-q1",
		Topic:      "GCP",
		Difficulty: "medium",
		Text:       "What is the default VPC network mode created automatically in a new GCP project?",
		Options: []Option{
			{ID: "opt1", Text: "Auto mode VPC", IsCorrect: true, Explanation: "Correct. Pre-creates subnets in each region."},
			{ID: "opt2", Text: "Custom mode VPC", IsCorrect: false, Explanation: "Requires manual subnet creation."},
			{ID: "opt3", Text: "Legacy network", IsCorrect: false, Explanation: "Legacy is deprecated."},
			{ID: "opt4", Text: "Shared VPC", IsCorrect: false, Explanation: "Shared VPC connects projects."},
		},
	},
	{
		ID:         "gcp-m-q2",
		Topic:      "GCP",
		Difficulty: "medium",
		Text:       "Which Google Kubernetes Engine (GKE) feature automatically upgrades cluster nodes?",
		Options: []Option{
			{ID: "opt1", Text: "Node auto-upgrade", IsCorrect: true, Explanation: "Correct. Keeps nodes at supported version."},
			{ID: "opt2", Text: "Cluster autoscaler", IsCorrect: false, Explanation: "Adds/removes nodes for workload demand."},
			{ID: "opt3", Text: "Workload Identity", IsCorrect: false, Explanation: "Maps K8s service accounts to IAM."},
			{ID: "opt4", Text: "Cloud Build triggers", IsCorrect: false, Explanation: "CI/CD service."},
		},
	},
	{
		ID:         "gcp-h-q1",
		Topic:      "GCP",
		Difficulty: "hard",
		Text:       "BigQuery uses which underlying storage format for columnar storage?",
		Options: []Option{
			{ID: "opt1", Text: "Capacitor", IsCorrect: true, Explanation: "Correct. Proprietary columnar format."},
			{ID: "opt2", Text: "Parquet", IsCorrect: false, Explanation: "External tables may use Parquet but internal uses Capacitor."},
			{ID: "opt3", Text: "ORC", IsCorrect: false, Explanation: "ORC not used natively."},
			{ID: "opt4", Text: "Avro", IsCorrect: false, Explanation: "Avro external."},
		},
	},
	{
		ID:         "gcp-h-q2",
		Topic:      "GCP",
		Difficulty: "hard",
		Text:       "Which service provides serverless analytics for real-time event ingestion similar to Kinesis?",
		Options: []Option{
			{ID: "opt1", Text: "Cloud Pub/Sub", IsCorrect: true, Explanation: "Correct. Pub/Sub streams events."},
			{ID: "opt2", Text: "Dataflow", IsCorrect: false, Explanation: "Dataflow processes streaming data but Pub/Sub ingests."},
			{ID: "opt3", Text: "Dataproc", IsCorrect: false, Explanation: "Dataproc runs Hadoop/Spark clusters."},
			{ID: "opt4", Text: "Composer", IsCorrect: false, Explanation: "Composer is managed Airflow."},
		},
	},
}
