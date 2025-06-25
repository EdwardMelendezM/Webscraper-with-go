import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WorkflowViewer } from './workflow-viewer';

describe('WorkflowViewer', () => {
  let component: WorkflowViewer;
  let fixture: ComponentFixture<WorkflowViewer>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [WorkflowViewer]
    })
    .compileComponents();

    fixture = TestBed.createComponent(WorkflowViewer);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
