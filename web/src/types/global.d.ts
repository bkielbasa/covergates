declare var VUE_BASE: any; // eslint-disable-line

declare interface User {
  login?: string;
  email?: string;
  avatar?: string;
}

declare interface Repository {
  URL: string;
  SCM: string;
  ReportID?: string;
  NameSpace: string;
  Name: string;
  Branch: string;
}

declare interface RepositorySetting {
  filters?: string[];
  mergePR?: boolean;
}

declare interface StatementHit {
  LineNumber: number;
  Hits: number;
}

declare interface SourceFile {
  Name: string;
  StatementCoverage: number;
  StatementHits: StatementHit[];
}

declare interface Coverage {
  Files?: SourceFile[];
  StatementCoverage: number | 0;
}

declare interface Report {
  branch?: string;
  commit: string;
  coverage?: Coverage;
  reportID: string;
  tag?: string;
  type?: string;
  files?: string[];
  createdAt?: string;
}
