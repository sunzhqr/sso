INSERT INTO apps
VALUES (1, 'test', 'test-secret')
ON CONFLICT DO NOTHING;