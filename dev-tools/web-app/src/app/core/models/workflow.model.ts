// Represents the data state at a single processing step
export interface WorkflowStep {
  stepName: 'Web Scraping' | 'Data Cleaning' | 'Lemmatization' | 'Data Enrichment';
  status: 'Completed' | 'In Progress' | 'Failed' | 'Pending';
  timestamp: Date;
  data: any; // Use a more specific type like string or Record<string, any>
  history?: {
    input: any;
    output: any;
    diff?: any; // To show what changed
  };
}

// Represents a complete job with all its steps
export interface ScrapingJob {
  id: string;
  sourceUrl: string;
  createdAt: Date;
  overallStatus: 'Completed' | 'In Progress' | 'Failed';
  workflowSteps: WorkflowStep[];
}
