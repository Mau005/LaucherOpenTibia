export namespace main {
	
	export class Players {
	    Online: number;
	    Max: number;
	    Peak: number;
	
	    static createFrom(source: any = {}) {
	        return new Players(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Online = source["Online"];
	        this.Max = source["Max"];
	        this.Peak = source["Peak"];
	    }
	}
	export class ServerInfo {
	    Uptime: string;
	    IP: string;
	    ServerName: string;
	    Port: string;
	    Location: string;
	    URL: string;
	    Server: string;
	    Version: string;
	    Client: string;
	
	    static createFrom(source: any = {}) {
	        return new ServerInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Uptime = source["Uptime"];
	        this.IP = source["IP"];
	        this.ServerName = source["ServerName"];
	        this.Port = source["Port"];
	        this.Location = source["Location"];
	        this.URL = source["URL"];
	        this.Server = source["Server"];
	        this.Version = source["Version"];
	        this.Client = source["Client"];
	    }
	}
	export class InfoResponse {
	    Version: string;
	    ServerInfo: ServerInfo;
	    Players: Players;
	
	    static createFrom(source: any = {}) {
	        return new InfoResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = source["Version"];
	        this.ServerInfo = this.convertValues(source["ServerInfo"], ServerInfo);
	        this.Players = this.convertValues(source["Players"], Players);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class NewsShort {
	    // Go type: time
	    CreatedAt: any;
	    IconID: number;
	    Description: string;
	
	    static createFrom(source: any = {}) {
	        return new NewsShort(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.CreatedAt = this.convertValues(source["CreatedAt"], null);
	        this.IconID = source["IconID"];
	        this.Description = source["Description"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

