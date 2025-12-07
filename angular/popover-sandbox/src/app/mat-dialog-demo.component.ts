import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatTooltipModule } from '@angular/material/tooltip';

@Component({
  selector: 'app-mat-dialog-demo',
  imports: [MatButtonModule, MatDialogModule, MatTooltipModule],
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <section class="dialog-demo">
      <h2 class="dialog-demo__title">Material Dialog の最小例</h2>
      <button mat-raised-button color="primary" type="button" (click)="openDialog()">
        ダイアログを開く
      </button>
      <button mat-stroked-button type="button" matTooltip="ツールチップのサンプル">
        ツールチップを見る
      </button>
    </section>
  `,
  styles: [
    `
      .dialog-demo {
        display: flex;
        flex-direction: column;
        gap: 12px;
        max-width: 360px;
      }

      .dialog-demo__title {
        margin: 0;
        font-size: 1.25rem;
      }
    `,
  ],
})
export class MatDialogDemoComponent {
  private readonly dialog = inject(MatDialog);

  openDialog(): void {
    this.dialog.open(BasicDialogContent, {
      width: '320px',
    });
  }
}

@Component({
  selector: 'app-basic-dialog-content',
  imports: [MatButtonModule, MatDialogModule],
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <h2 mat-dialog-title>こんにちは 👋</h2>
    <div mat-dialog-content>Angular Material のダイアログの最小サンプルです。</div>
    <div mat-dialog-actions align="end">
      <button mat-stroked-button type="button" (click)="close()">閉じる</button>
    </div>
  `,
})
export class BasicDialogContent {
  private readonly dialogRef = inject(MatDialogRef<BasicDialogContent>);

  close(): void {
    this.dialogRef.close();
  }
}
