import { OverlayRef, PositionStrategy } from '@angular/cdk/overlay';
import { MatTooltip } from '@angular/material/tooltip';

type PopoverCapableStrategy = PositionStrategy & {
  withPopoverLocation?(location: 'inline' | 'global' | {type: 'parent'; element: Element}): void;
  getPopoverInsertionPoint?(): Element | null | {type: 'parent'; element: Element};
};

/** MatTooltip がグローバル挿入に固定している popover 位置を inline に差し替えるパッチ。 */
export function patchMatTooltipPopoverInline(): void {
  const tooltipProto = MatTooltip.prototype as unknown as {
    _createOverlay?: (...args: unknown[]) => OverlayRef;
    __popoverPatched?: boolean;
  };

  if (!tooltipProto._createOverlay || tooltipProto.__popoverPatched) {
    return;
  }

  const originalCreateOverlay = tooltipProto._createOverlay;

  tooltipProto._createOverlay = function patchedCreateOverlay(
    this: MatTooltip,
    ...args: unknown[]
  ) {
    const overlayRef = originalCreateOverlay.apply(this, args) as OverlayRef;
    const strategy = overlayRef.getConfig().positionStrategy as PopoverCapableStrategy | undefined;

    // MatTooltip は withPopoverLocation('global') を固定で呼ぶため、ここで inline に上書きし、
    // 必要ならホストを挿入し直す。
    strategy?.withPopoverLocation?.('inline');

    const insertionPoint = strategy?.getPopoverInsertionPoint?.();
    const host = overlayRef.hostElement;

    if (host && insertionPoint) {
      if (insertionPoint instanceof Element) {
        insertionPoint.after(host);
      } else if (insertionPoint.type === 'parent' && insertionPoint.element instanceof Element) {
        insertionPoint.element.appendChild(host);
      }
    }

    return overlayRef;
  };

  tooltipProto.__popoverPatched = true;
}
