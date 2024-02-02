export namespace controllers {
	
	export class ALLNoteDTO {
	    Data: number;
	    ParentUid?: string;
	    UserUid: string;
	    SubjectOld?: string;
	    CategoryUid: number;
	    SubCategory: number;
	    Note: string;
	    Note_Plain: string;
	    Priority: number;
	    CommunicationCode?: number;
	    IsReminder: boolean;
	    CompletedBy?: string;
	    // Go type: time
	    CompletedDate?: any;
	    // Go type: time
	    ScheduledStart?: any;
	    // Go type: time
	    ScheduledEnd?: any;
	    OldRecId?: string;
	    // Go type: time
	    AppTimeStamp?: any;
	    AppLastUpdatedBy?: string;
	    AppCreatedBy?: string;
	    ActivityUid?: string;
	    Subject?: string;
	    ContactName?: string;
	    AppTimeStampYear: number;
	    SysTimeStamp: number[];
	    Note_PlainResume?: string;
	    ParentDepartmentUid: number;
	    ParentAccount: string;
	    ParentType?: number;
	    // Go type: time
	    AppLastUpdated?: any;
	
	    static createFrom(source: any = {}) {
	        return new ALLNoteDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Data = source["Data"];
	        this.ParentUid = source["ParentUid"];
	        this.UserUid = source["UserUid"];
	        this.SubjectOld = source["SubjectOld"];
	        this.CategoryUid = source["CategoryUid"];
	        this.SubCategory = source["SubCategory"];
	        this.Note = source["Note"];
	        this.Note_Plain = source["Note_Plain"];
	        this.Priority = source["Priority"];
	        this.CommunicationCode = source["CommunicationCode"];
	        this.IsReminder = source["IsReminder"];
	        this.CompletedBy = source["CompletedBy"];
	        this.CompletedDate = this.convertValues(source["CompletedDate"], null);
	        this.ScheduledStart = this.convertValues(source["ScheduledStart"], null);
	        this.ScheduledEnd = this.convertValues(source["ScheduledEnd"], null);
	        this.OldRecId = source["OldRecId"];
	        this.AppTimeStamp = this.convertValues(source["AppTimeStamp"], null);
	        this.AppLastUpdatedBy = source["AppLastUpdatedBy"];
	        this.AppCreatedBy = source["AppCreatedBy"];
	        this.ActivityUid = source["ActivityUid"];
	        this.Subject = source["Subject"];
	        this.ContactName = source["ContactName"];
	        this.AppTimeStampYear = source["AppTimeStampYear"];
	        this.SysTimeStamp = source["SysTimeStamp"];
	        this.Note_PlainResume = source["Note_PlainResume"];
	        this.ParentDepartmentUid = source["ParentDepartmentUid"];
	        this.ParentAccount = source["ParentAccount"];
	        this.ParentType = source["ParentType"];
	        this.AppLastUpdated = this.convertValues(source["AppLastUpdated"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class DataSourceResponseDTO {
	    Data: ALLNoteDTO[];
	    Total: number;
	
	    static createFrom(source: any = {}) {
	        return new DataSourceResponseDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Data = this.convertValues(source["Data"], ALLNoteDTO);
	        this.Total = source["Total"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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

