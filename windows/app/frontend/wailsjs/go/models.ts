export namespace main {
	
	export class AppInfo {
	    id: string;
	    deviceToken: string;
	    name: string;
	    version: string;
	    build: string;
	    tag: string;
	    buildType: string;
	    company: string;
	    website: string;
	    endpoint: string;
	    description: string;
	
	    static createFrom(source: any = {}) {
	        return new AppInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.deviceToken = source["deviceToken"];
	        this.name = source["name"];
	        this.version = source["version"];
	        this.build = source["build"];
	        this.tag = source["tag"];
	        this.buildType = source["buildType"];
	        this.company = source["company"];
	        this.website = source["website"];
	        this.endpoint = source["endpoint"];
	        this.description = source["description"];
	    }
	}
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
	
	    static createFrom(source: any = {}) {
	        return new DeviceInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hostname = source["hostname"];
	        this.domain = source["domain"];
	        this.user = source["user"];
	        this.os = source["os"];
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
	export class SecurityInfo {
	    windowsDefender: boolean;
	    firewall: boolean;
	    tpm: boolean;
	    bitLocker: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SecurityInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.windowsDefender = source["windowsDefender"];
	        this.firewall = source["firewall"];
	        this.tpm = source["tpm"];
	        this.bitLocker = source["bitLocker"];
	    }
	}

}

