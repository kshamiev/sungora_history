describe('securityRetryQueue', function () {
  var queue = 0;

  function mockRetryItem () {
    return jasmine.createSpyObj('retryItem', ['cancel', 'retry']);
  }

  beforeEach(
    module('security.retryQueue')
  );

  beforeEach(
    inject(
      function (_securityRetryQueue_) {
        queue = _securityRetryQueue_;
      }
    )
  );

  describe('hasMore', function () {
    it('initialy has no items to retry', function () {
      expect(queue.hasMore).toBeDefined();
      expect(queue.hasMore()).toBe(false);
    });

    it('has more items ones one item was pushes', function () {
      queue.push(mockRetryItem());
      expect(queue.hasMore()).toBe(true);
    });
  });

  describe('pushRetryFn', function () {
    it('adds a new item to the queue', function () {
     queue.pushRetryFn(function () {});
     expect(queue.hasMore()).toBe(true);
    });

    it('adds a reason to the retry', function () {
      var reason = 'some_reason';
      queue.pushRetryFn(reason, function () {});
      expect(queue.retryReason()).toBe(reason);
    });

    it('does not add a reason to the retry if not specified', function () {
      queue.pushRetryFn(function () {});
      expect(queue.retryReason()).not.toBeDefined();
    });
  });

  describe('retryAll', function () {
    it('should not fail if the queue is empty', function () {
      queue.retryAll(function (item) {});
      expect(queue.hasMore()).toBe(false);
    });
    
    it('should empty the queue', function () {
      queue.push(mockRetryItem());
      queue.push(mockRetryItem());
      queue.push(mockRetryItem());
      expect(queue.hasMore()).toBe(true);
      queue.retryAll(function (item) {});
      expect(queue.hasMore()).toBe(false);
    });
  });

  describe('cancelAll', function () {
    it('should empty the queye', function () {
      queue.push(mockRetryItem());
      queue.push(mockRetryItem());
      queue.push(mockRetryItem());
      expect(queue.hasMore()).toBe(true);
      queue.cancelAll(function (item) {});
      expect(queue.hasMore()).toBe(false);
    });
  });
});
