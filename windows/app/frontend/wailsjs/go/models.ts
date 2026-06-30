export namespace main {
	
	export class CpuFullDetails {
	    model: string;
	    physicalCores: number;
	    logicalThreads: number;
	    maxSpeed: string;
	    currentSpeed: string;
	    l2Cache: string;
	    l3Cache: string;
	    liveUsage: string;
	
	    static createFrom(source: any = {}) {
	        return new CpuFullDetails(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.model = source["model"];
	        this.physicalCores = source["physicalCores"];
	        this.logicalThreads = source["logicalThreads"];
	        this.maxSpeed = source["maxSpeed"];
	        this.currentSpeed = source["currentSpeed"];
	        this.l2Cache = source["l2Cache"];
	        this.l3Cache = source["l3Cache"];
	        this.liveUsage = source["liveUsage"];
	    }
	}
	export class DeviceInfo {
	    hostname: string;
	    domain: string;
	    user: string;
	    os: string;
	    agentVersion: string;
	    windowsDefender: boolean;
	    firewall: boolean;
	    tpm: boolean;
	    bitLocker: boolean;
	
	    static createFrom(source: any = {}) {
	        return new DeviceInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hostname = source["hostname"];
	        this.domain = source["domain"];
	        this.user = source["user"];
	        this.os = source["os"];
	        this.agentVersion = source["agentVersion"];
	        this.windowsDefender = source["windowsDefender"];
	        this.firewall = source["firewall"];
	        this.tpm = source["tpm"];
	        this.bitLocker = source["bitLocker"];
	    }
	}
	export class MemoryDetails {
	    used: string;
	    capacity: string;
	    totalslots: number;
	    slotused: number;
	    free: string;
	
	    static createFrom(source: any = {}) {
	        return new MemoryDetails(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.used = source["used"];
	        this.capacity = source["capacity"];
	        this.totalslots = source["totalslots"];
	        this.slotused = source["slotused"];
	        this.free = source["free"];
	    }
	}

}

