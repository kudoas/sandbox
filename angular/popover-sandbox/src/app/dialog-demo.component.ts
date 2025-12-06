import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-dialog-demo',
  imports: [MatButtonModule, MatDialogModule],
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <section class="dialog-demo">
      <h2 class="dialog-demo__title">Angular Material Dialog デモ</h2>
      <button mat-raised-button color="primary" type="button" (click)="openDialog()">
        ダイアログを開く
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
export class DialogDemoComponent {
  private readonly dialog = inject(MatDialog);

  openDialog(): void {
    this.dialog.open(SimpleDialogComponent, {
      width: '320px',
    });
  }
}

@Component({
  selector: 'app-simple-dialog',
  imports: [MatButtonModule, MatDialogModule],
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <h2 mat-dialog-title>こんにちは 👋</h2>
    <div mat-dialog-content>Angular Material のダイアログの最小例です。</div>
    <div mat-dialog-actions align="end">
      <button mat-stroked-button type="button" (click)="close()">閉じる</button>
    </div>
  `,
})
export class SimpleDialogComponent {
  private readonly dialogRef = inject(MatDialogRef<SimpleDialogComponent>);

  close(): void {
    this.dialogRef.close();
  }
}
