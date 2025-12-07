import { ChangeDetectionStrategy, Component, signal } from '@angular/core';
import { OverlayModule } from '@angular/cdk/overlay';

@Component({
  selector: 'app-dialog-demo',
  imports: [OverlayModule],
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <section class="dialog-demo" cdkOverlayOrigin #trigger="cdkOverlayOrigin">
      <h2 class="dialog-demo__title">Inline Overlay Dialog の例</h2>
      <button mat-raised-button color="primary" type="button" (click)="toggle()">
        {{ isOpen() ? '閉じる' : 'ダイアログを開く' }}
      </button>
    </section>

    <ng-template
      cdkConnectedOverlay
      [cdkConnectedOverlayOrigin]="trigger"
      [cdkConnectedOverlayOpen]="isOpen()"
      [cdkConnectedOverlayHasBackdrop]="true"
      [cdkConnectedOverlayBackdropClass]="'cdk-overlay-transparent-backdrop'"
      [cdkConnectedOverlayUsePopover]="'inline'"
      (backdropClick)="close()"
    >
      <div
        class="popover-panel"
        role="dialog"
        aria-modal="true"
        aria-label="サンプルダイアログ"
        tabindex="0"
      >
        <h3 class="popover-panel__title">こんにちは 👋</h3>
        <p class="popover-panel__body">Inline overlay で描画しています。</p>
        <div class="popover-panel__actions">
          <button mat-stroked-button type="button" (click)="close()">閉じる</button>
        </div>
      </div>
    </ng-template>
  `,
  styles: [
    `
      .dialog-demo {
        display: flex;
        flex-direction: column;
        gap: 12px;
        max-width: 440px;
        margin: 32px auto 0;
        padding: 16px;
        border: 1px solid #d5d7da;
        border-radius: 8px;
        background: #fafbfd;
      }

      .dialog-demo__title {
        margin: 0;
        font-size: 1.25rem;
      }

      .dialog-demo__hint {
        margin: 0;
        color: #4b5563;
      }

      .popover-panel {
        display: flex;
        flex-direction: column;
        gap: 12px;
        min-width: 260px;
        padding: 16px;
        border-radius: 10px;
        background: #ffffff;
        box-shadow: 0 10px 30px rgba(0, 0, 0, 0.12), 0 2px 8px rgba(0, 0, 0, 0.08);
        outline: none;
      }

      .popover-panel__title {
        margin: 0;
        font-size: 1.1rem;
      }

      .popover-panel__body {
        margin: 0;
        color: #374151;
      }

      .popover-panel__actions {
        display: flex;
        justify-content: flex-end;
      }
    `,
  ],
})
export class DialogDemoComponent {
  protected readonly isOpen = signal(false);

  toggle(): void {
    this.isOpen.update((open) => !open);
  }

  close(): void {
    this.isOpen.set(false);
  }
}
