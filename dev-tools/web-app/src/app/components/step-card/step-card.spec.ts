import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StepCard } from './step-card';

describe('StepCard', () => {
  let component: StepCard;
  let fixture: ComponentFixture<StepCard>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [StepCard]
    })
    .compileComponents();

    fixture = TestBed.createComponent(StepCard);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
